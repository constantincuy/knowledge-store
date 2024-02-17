package worker

import (
	"context"
	"github.com/constantincuy/knowledgestore/internal/core/domain/file"
	"github.com/constantincuy/knowledgestore/internal/core/domain/knowledgebase"
	"github.com/constantincuy/knowledgestore/internal/ports"
	"github.com/google/uuid"
	"github.com/vladopajic/go-actor/actor"
	"log"
	"os"
	"path"
	"time"
)

type FileWorker struct {
	name         string
	mailbox      actor.MailboxSender[file.Downloaded]
	storage      ports.Storage
	fileRepo     ports.FileRepo
	downloadPath string
}

type IndexWorkType int

var (
	IndexCreate IndexWorkType = 1
	IndexUpdate IndexWorkType = 2
)

func indexTypeToString(it IndexWorkType) string {
	if it == IndexCreate {
		return "Created"
	}

	if it == IndexUpdate {
		return "Updated"
	}

	return ""
}

func (w *FileWorker) DoWork(ctx actor.Context) actor.WorkerStatus {
	select {
	case <-ctx.Done():
		return actor.WorkerEnd

	case <-time.After(time.Second * 10):
		name, err := knowledgebase.NewName(w.name)
		if err != nil {
			log.Printf("Failed to start knowledge base %s file worker: %s", w.name, err)
			return actor.WorkerEnd
		}
		provider, _ := file.NewProvider(w.storage.Provider())
		files, _ := w.fileRepo.GetAllProviderFiles(ctx, name, provider)
		list, _ := file.NewList(files)
		changeList, _ := w.storage.GetChangedDocuments(ctx, file.NewFilesystem(list))
		for _, meta := range changeList.Deleted {
			err := w.fileRepo.Delete(ctx, name, provider, meta.Path)
			if err != nil {
				log.Printf("[%s] Failed to delete file index %s\n", w.name, err)
				return actor.WorkerContinue
			}
			log.Printf("[%s] Deleted file index %s\n", w.name, meta.Path)
		}

		w.sendFilesToIndexingWorker(ctx, name, IndexCreate, changeList.Created)
		w.sendFilesToIndexingWorker(ctx, name, IndexUpdate, changeList.Updated)

		return actor.WorkerContinue
	}
}

func (w *FileWorker) sendFilesToIndexingWorker(ctx context.Context, name knowledgebase.Name, action IndexWorkType, list file.List) {
	for _, meta := range list {
		var err error

		switch action {
		case IndexCreate:
			err = w.fileRepo.Add(ctx, name, meta)
		case IndexUpdate:
			err = w.fileRepo.Update(ctx, name, meta)
		}

		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("[%s] %s file index %s\n", w.name, indexTypeToString(action), meta.Path)
		downloadFile := path.Join(w.downloadPath, w.name+"_"+uuid.UUID(meta.Id).String()+".txt")
		f, err := os.Create(downloadFile)
		if err != nil {
			log.Println(err)
			return
		}
		w.storage.DownloadDocument(ctx, meta.Path, f)
		err = f.Close()
		if err != nil {
			log.Println(err)
		}
		err = w.mailbox.Send(ctx, file.NewDownloaded(downloadFile, meta))
		if err != nil {
			log.Println(err)
		}
	}
}

func NewFileWorker(name string, fileRepo ports.FileRepo, storage ports.Storage, mailbox actor.MailboxSender[file.Downloaded]) FileWorker {
	dir, _ := os.Getwd()
	downloadPath := path.Join(dir, "download", "fake", storage.Provider())
	_ = os.MkdirAll(downloadPath, os.ModePerm)
	return FileWorker{name, mailbox, storage, fileRepo, downloadPath}
}
