package actor

import (
	"context"
	"github.com/constantincuy/knowledgestore/internal/core/domain/file"
	"github.com/constantincuy/knowledgestore/internal/core/service/documents"
	"github.com/constantincuy/knowledgestore/internal/core/worker"
	"github.com/constantincuy/knowledgestore/internal/ports"
	"github.com/vladopajic/go-actor/actor"
	"log"
)

type KnowledgeBase struct {
	name  string
	actor actor.Actor
}

func (k KnowledgeBase) Start() {
	k.actor.Start()
}

func (k KnowledgeBase) Stop() {
	k.actor.Stop()
}

func NewKnowledgeBase(name string, fileRepo ports.FileRepo, docService documents.Api, embedding ports.EmbeddingExtractor, storage ports.Storage) KnowledgeBase {
	mailbox := actor.NewMailbox[file.Downloaded]()

	fw := worker.NewFileWorker(name, fileRepo, storage, mailbox)
	dw := worker.NewDocumentWorker(name, fileRepo, docService, embedding, mailbox)
	afw := actor.New(&fw)
	adw := actor.New(&dw)
	return KnowledgeBase{name, actor.Combine(mailbox, afw, adw).WithOptions(actor.OptOnStartCombined(func(ctx context.Context) { log.Printf("Started actor for knowledgebase '%s'!", name) })).Build()}
}
