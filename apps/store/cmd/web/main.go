package main

import (
	"github.com/constantincuy/knowledgestore/internal/adapters/primary/rest"
	"github.com/constantincuy/knowledgestore/internal/adapters/secondary/embedding/localembed"
	"github.com/constantincuy/knowledgestore/internal/adapters/secondary/postgres"
	"github.com/constantincuy/knowledgestore/internal/adapters/secondary/storage/fake"
	"github.com/constantincuy/knowledgestore/internal/core/domain/file"
	"github.com/constantincuy/knowledgestore/internal/core/service/documents"
	"github.com/constantincuy/knowledgestore/internal/core/service/knowledgebases"
	"github.com/constantincuy/knowledgestore/internal/core/worker"
	_ "github.com/lib/pq"
	"github.com/vladopajic/go-actor/actor"
	"log"
	"sync"
)

func main() {
	pr := postgres.NewProvider("root", "root", "localhost", "disable")
	knowledgeBaseRepo, err := postgres.NewKnowledgeBaseRepo(pr)
	fileRepo, err := postgres.NewFileRepo(pr)
	docRepo, err := postgres.NewDocumentRepo(pr)

	if err != nil {
		log.Fatal("failed to init postgres repo: %w", err)
	}

	knowledgeService := knowledgebases.NewService(knowledgeBaseRepo)
	docService := documents.NewService(docRepo)
	embedding := localembed.Extractor{}

	server := rest.New(knowledgeService)

	fakeStorage := fake.NewStorage()

	mb := actor.NewMailbox[file.Downloaded]()
	fw := worker.NewFileWorker(fileRepo, fakeStorage, mb)
	dw := worker.NewDocumentWorker(fileRepo, docService, embedding, mb)
	afw := actor.New(&fw)
	adw := actor.New(&dw)
	a := actor.Combine(mb, afw, adw).Build()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go server.Run()
	wg.Add(1)
	go a.Start()

	wg.Wait()
}
