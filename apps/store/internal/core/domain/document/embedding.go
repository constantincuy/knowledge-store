package document

import (
	"errors"
)

var (
	ErrEmptyEmbedding = errors.New("empty embedding supplied")
)

type Embedding []float32

func NewEmbedding(un []float32) (Embedding, error) {
	if len(un) == 0 {
		return make([]float32, 0), ErrEmptyEmbedding
	}

	return Embedding(un), nil
}
