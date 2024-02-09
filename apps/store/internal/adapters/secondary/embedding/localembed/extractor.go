package localembed

import (
	"bytes"
	"encoding/json"
	"github.com/constantincuy/knowledgestore/internal/adapters/secondary/embedding"
	"github.com/constantincuy/knowledgestore/internal/ports"
	"io"
	"net/http"
	"time"
)

type embeddingGeneratorResult struct {
	Dims    []int     `json:"dims"`
	Type    string    `json:"type"`
	Vectors []float32 `json:"vectors"`
}

type Extractor struct{}

func (e Extractor) Extract(strings []string) (embedding.Embedding, error) {
	b, err := json.Marshal(strings)
	if err != nil {
		return embedding.Embedding{}, ports.ErrGeneratingEmbeddings
	}
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Post("http://localhost:3000/embeddings/generate", "application/json", bytes.NewReader(b))
	if err != nil {
		return embedding.Embedding{}, ports.ErrGeneratingEmbeddings
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	var result embeddingGeneratorResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return embedding.Embedding{}, ports.ErrGeneratingEmbeddings
	}

	return embedding.Embedding{Vectors: result.Vectors}, nil
}
