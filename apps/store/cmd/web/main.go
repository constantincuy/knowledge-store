package main

import (
	"github.com/constantincuy/knowledgestore/internal/adapters/primary/rest"
	"github.com/constantincuy/knowledgestore/internal/adapters/secondary/postgres"
	"github.com/constantincuy/knowledgestore/internal/core/service/knowledgebases"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

func main() {
	pr := postgres.NewProvider("root", "root", "localhost", "disable")
	knowledgeBaseRepo, err := postgres.NewKnowledgeBaseRepo(pr)

	if err != nil {
		log.Fatal("failed to init postgres repo: %w", err)
	}

	knowledgeService := knowledgebases.NewService(knowledgeBaseRepo)

	server := rest.New(knowledgeService)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go server.Run()

	wg.Wait()
}
