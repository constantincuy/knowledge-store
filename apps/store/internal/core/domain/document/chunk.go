package document

import (
	"errors"
)

var (
	ErrInvalidChunk = errors.New("invalid chunk supplied")
)

type Chunk int

func NewChunk(number int) (Chunk, error) {
	if number < 1 {
		return 1, ErrInvalidChunk
	}

	return Chunk(number), nil
}
