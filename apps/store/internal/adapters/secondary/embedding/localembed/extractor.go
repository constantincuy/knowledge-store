package localembed

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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

type Extractor struct {
	host string
	port int
}

func (e Extractor) Extract(strings []string) (embedding.Embedding, error) {
	b, err := json.Marshal(strings)
	if err != nil {
		return embedding.Embedding{}, errors.Join(ports.ErrGeneratingEmbeddings, err)
	}
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Post(fmt.Sprintf("http://%s:%d/embeddings/generate", e.host, e.port), "application/json", bytes.NewReader(b))
	if err != nil {
		return embedding.Embedding{}, errors.Join(ports.ErrGeneratingEmbeddings, err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	var result embeddingGeneratorResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return embedding.Embedding{}, errors.Join(ports.ErrGeneratingEmbeddings, err)
	}

	return embedding.Embedding{Vectors: result.Vectors}, nil
}

func NewExtractor(host string, port int) Extractor {
	return Extractor{host, port}
}
