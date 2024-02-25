package worker

import (
	"bufio"
	"github.com/constantincuy/knowledgestore/internal/core/domain/document"
	"github.com/constantincuy/knowledgestore/internal/core/domain/file"
	"github.com/constantincuy/knowledgestore/internal/core/service/documents"
	"github.com/constantincuy/knowledgestore/internal/ports"
	"github.com/google/uuid"
	"github.com/vladopajic/go-actor/actor"
	"log"
	"os"
)

const WordsPerChunk = 256

type DocumentWorker struct {
	name       string
	mailbox    actor.MailboxReceiver[file.Downloaded]
	fileRepo   ports.FileRepo
	embedding  ports.EmbeddingExtractor
	docService documents.Api
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
		defer f.Close()

		chunked, err := document.NewChunkedDocument(bufio.NewReader(f), WordsPerChunk)
		if err != nil {
			log.Println(err)
			return actor.WorkerContinue
		}

		w.docService.Delete(ctx, documents.DeleteDocumentReq{
			KnowledgeBase: w.name,
			FileId:        uuid.UUID(downloaded.Meta.Id),
		})
		for i, chunk := range chunked.Chunks() {
			data := make([]string, 1)
			data[0] = chunk

			embed, err := w.embedding.Extract(data)
			if err != nil {
				log.Println(err)
				return actor.WorkerContinue
			}

			_, _ = w.docService.Create(ctx, documents.AddDocumentReq{
				KnowledgeBase: w.name,
				FileId:        uuid.UUID(downloaded.Meta.Id),
				Chunk:         i + 1,
				Embedding:     embed.Vectors,
			})
		}
		log.Printf("[%s] Indexed %d chunks for file %s\n", w.name, len(chunked.Chunks()), downloaded.Meta.Path)
		f.Close()
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
