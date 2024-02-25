package document

import (
	"bufio"
	"io"
	"strings"
)

type Chunked struct {
	chunks    []string
	chunkSize int
}

func (c Chunked) GetChunk(chunk int) string {
	return c.chunks[chunk-1]
}

func (c Chunked) Chunks() []string {
	return c.chunks
}

func readChunks(file io.Reader, chunkSize int) ([]string, error) {
	chunks := make([]string, 0)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	words := 0
	chunk := ""
	for scanner.Scan() {
		word := scanner.Text()
		chunk += word + " "
		words++

		if words == chunkSize {
			chunks = append(chunks, strings.TrimSpace(chunk))
			chunk = ""
			words = 0
		}
	}

	if words < chunkSize {
		chunks = append(chunks, strings.TrimSpace(chunk))
		chunk = ""
		words = 0
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if chunk != "" {
		chunks = append(chunks, strings.TrimSpace(chunk))
	}

	return chunks, nil
}

func NewChunkedDocument(file io.Reader, chunkSize int) (Chunked, error) {
	chunks, err := readChunks(file, chunkSize)

	if err != nil {
		return Chunked{}, err
	}

	return Chunked{chunks, chunkSize}, nil
}
