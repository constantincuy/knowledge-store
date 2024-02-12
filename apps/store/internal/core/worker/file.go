package worker

import (
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

func (w *FileWorker) DoWork(ctx actor.Context) actor.WorkerStatus {
	select {
	case <-ctx.Done():
		return actor.WorkerEnd

	case <-time.After(time.Second * 10):
		name, _ := knowledgebase.NewName(w.name)
		provider, _ := file.NewProvider(w.storage.Provider())
		files, _ := w.fileRepo.GetAllProviderFiles(ctx, name, provider)
		list, _ := file.NewList(files)
		changeList, _ := w.storage.GetChangedDocuments(ctx, file.NewFilesystem(list))
		for _, meta := range changeList.Created {
			err := w.fileRepo.Add(ctx, name, meta)
			if err != nil {
				log.Println(err)
				return actor.WorkerContinue
			}
			log.Printf("Created file %s\n", meta.Path)
			downloadFile := path.Join(w.downloadPath, uuid.UUID(meta.Id).String()+".txt")
			f, err := os.Create(downloadFile)
			if err != nil {
				log.Println(err)
				return actor.WorkerContinue
			}
			w.storage.DownloadDocument(ctx, meta.Path, f)
			err = w.mailbox.Send(ctx, file.NewDownloaded(downloadFile, meta))
			if err != nil {
				log.Println(err)
			}
		}
		return actor.WorkerContinue
	}
}

func NewFileWorker(name string, fileRepo ports.FileRepo, storage ports.Storage, mailbox actor.MailboxSender[file.Downloaded]) FileWorker {
	dir, _ := os.Getwd()
	downloadPath := path.Join(dir, "download", "fake")
	_ = os.MkdirAll(downloadPath, os.ModePerm)
	return FileWorker{name, mailbox, storage, fileRepo, downloadPath}
}
