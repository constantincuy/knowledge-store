package main

import (
	"context"
	"github.com/constantincuy/knowledgestore/internal/adapters/primary/rest"
	"github.com/constantincuy/knowledgestore/internal/adapters/secondary/embedding/localembed"
	"github.com/constantincuy/knowledgestore/internal/adapters/secondary/postgres"
	"github.com/constantincuy/knowledgestore/internal/adapters/secondary/storage/fake"
	"github.com/constantincuy/knowledgestore/internal/core/actor"
	"github.com/constantincuy/knowledgestore/internal/core/service/documents"
	"github.com/constantincuy/knowledgestore/internal/core/service/files"
	"github.com/constantincuy/knowledgestore/internal/core/service/knowledgebases"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

func main() {
	pr := postgres.NewProvider("root", "root", "localhost", "disable")
	postgres.RunSetup(context.Background(), pr)
	knowledgeBaseRepo, err := postgres.NewKnowledgeBaseRepo(pr)
	fileRepo, err := postgres.NewFileRepo(pr)
	docRepo, err := postgres.NewDocumentRepo(pr)

	if err != nil {
		log.Fatalf("failed to init postgres repo: %s", err)
	}

	embedding := localembed.Extractor{}
	knowledgeService := knowledgebases.NewService(knowledgeBaseRepo)
	docService := documents.NewService(docRepo)
	fileService := files.NewService(fileRepo, embedding)

	fakeStorage := fake.NewStorage()
	mn := actor.NewManager(knowledgeBaseRepo, fileRepo, docService, embedding, fakeStorage)

	server := rest.New(&mn, knowledgeService, fileService)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go server.Run()
	err = mn.Hydrate()
	if err != nil {
		log.Fatalf("Could not hydrate worker threads: %s", err)
	}

	wg.Wait()
}
