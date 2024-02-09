package ports

import (
	"errors"
	"github.com/constantincuy/knowledgestore/internal/adapters/secondary/embedding"
)

var (
	ErrGeneratingEmbeddings = errors.New("error generating embedding")
)

type EmbeddingExtractor interface {
	Extract([]string) (embedding.Embedding, error)
}
