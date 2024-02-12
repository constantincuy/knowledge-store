package worker

import (
	"bufio"
	"github.com/constantincuy/knowledgestore/internal/core/domain/file"
	"github.com/constantincuy/knowledgestore/internal/core/service/documents"
	"github.com/constantincuy/knowledgestore/internal/ports"
	"github.com/google/uuid"
	"github.com/vladopajic/go-actor/actor"
	"io"
	"log"
	"os"
	"strings"
)

type DocumentWorker struct {
	name       string
	mailbox    actor.MailboxReceiver[file.Downloaded]
	fileRepo   ports.FileRepo
	embedding  ports.EmbeddingExtractor
	docService documents.Api
}

func readChunks(file io.Reader) ([]string, error) {
	chunks := make([]string, 0)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	words := 0
	chunk := ""
	for scanner.Scan() {
		word := scanner.Text()
		chunk += word + " "
		words++

		if words == 6000 {
			chunks = append(chunks, strings.TrimSpace(chunk))
			chunk = ""
			words = 0
		}
	}

	if words < 6000 {
		chunks = append(chunks, strings.TrimSpace(chunk))
		chunk = ""
		words = 0
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if chunk != "" {
		chunks = append(chunks, strings.TrimSpace(chunk))
	}

	return chunks, nil
}

func (w *DocumentWorker) DoWork(ctx actor.Context) actor.WorkerStatus {
	select {
	case <-ctx.Done():
		return actor.WorkerEnd

	case downloaded := <-w.mailbox.ReceiveC():
		f, err := os.Open(downloaded.DownloadPath)
		if err != nil {
			log.Println(err)
			return actor.WorkerContinue
		}

		chunks, err := readChunks(bufio.NewReader(f))
		if err != nil {
			log.Println(err)
			return actor.WorkerContinue
		}

		for i, chunk := range chunks {
			data := make([]string, 1)
			data[0] = chunk

			embed, err := w.embedding.Extract(data)
			if err != nil {
				log.Println(err)
				return actor.WorkerContinue
			}

			log.Printf("[%s] Indexed chunk %d of file %s\n", w.name, i+1, downloaded.Meta.Path)
			_, _ = w.docService.Create(ctx, documents.AddDocumentReq{
				KnowledgeBase: w.name,
				FileId:        uuid.UUID(downloaded.Meta.Id),
				Chunk:         i + 1,
				Embedding:     embed.Vectors,
			})
		}

		err = os.Remove(downloaded.DownloadPath)
		if err != nil {
			log.Println(err)
		}

		return actor.WorkerContinue
	}
}

func NewDocumentWorker(name string, fileRepo ports.FileRepo, docService documents.Api, embedding ports.EmbeddingExtractor, mailbox actor.MailboxReceiver[file.Downloaded]) DocumentWorker {
	return DocumentWorker{
		name:       name,
		mailbox:    mailbox,
		fileRepo:   fileRepo,
		embedding:  embedding,
		docService: docService,
	}
}
