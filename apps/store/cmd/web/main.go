package main

import (
	"context"
	"github.com/constantincuy/knowledgestore/internal/adapters/primary/rest"
	"github.com/constantincuy/knowledgestore/internal/adapters/secondary/embedding/localembed"
	"github.com/constantincuy/knowledgestore/internal/adapters/secondary/postgres"
	"github.com/constantincuy/knowledgestore/internal/adapters/secondary/storage/zip"
	"github.com/constantincuy/knowledgestore/internal/core/actor"
	"github.com/constantincuy/knowledgestore/internal/core/service/documents"
	"github.com/constantincuy/knowledgestore/internal/core/service/files"
	"github.com/constantincuy/knowledgestore/internal/core/service/knowledgebases"
	_ "github.com/lib/pq"
	"github.com/sethvargo/go-envconfig"
	"log"
	"os"
	"path"
	"sync"
)

type AppConfig struct {
	PostgresHost       string `env:"EMBD_STORE_DB_HOST, required"`
	PostgresUser       string `env:"EMBD_STORE_DB_USER, required"`
	PostgresPassword   string `env:"EMBD_STORE_DB_PASSWORD, required"`
	LocalEmbeddingHost string `env:"EMBD_STORE_EMBEDDING_EXTRACTOR_HOST, required"`
	LocalEmbeddingPort int    `env:"EMBD_STORE_EMBEDDING_EXTRACTOR_PORT, required"`
}

func main() {
	var cfg AppConfig
	if err := envconfig.Process(context.Background(), &cfg); err != nil {
		log.Fatal(err)
	}

	pr := postgres.NewProvider(cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresHost, "disable")
	postgres.RunSetup(context.Background(), pr)
	knowledgeBaseRepo, err := postgres.NewKnowledgeBaseRepo(pr)
	fileRepo, err := postgres.NewFileRepo(pr)
	docRepo, err := postgres.NewDocumentRepo(pr)

	if err != nil {
		log.Fatalf("failed to init postgres repo: %s", err)
	}

	embedding := localembed.NewExtractor(cfg.LocalEmbeddingHost, cfg.LocalEmbeddingPort)
	knowledgeService := knowledgebases.NewService(knowledgeBaseRepo)
	docService := documents.NewService(docRepo)
	fileService := files.NewService(fileRepo, embedding)

	dir, _ := os.Getwd()
	zipPath := path.Join(dir, "example", "example.zip")
	fakeStorage := zip.NewStorage(zipPath)
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
