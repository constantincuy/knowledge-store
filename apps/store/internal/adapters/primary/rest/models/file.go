package models

import (
	"github.com/constantincuy/knowledgestore/internal/core/service/files"
	"github.com/google/uuid"
	"time"
)

type File struct {
	Meta    FileMeta `json:"meta"`
	Content []string `json:"content"`
}

type FileMeta struct {
	Id       string     `json:"id"`
	Path     string     `json:"path"`
	Provider string     `json:"provider"`
	Chunks   []Document `json:"chunks"`
	Created  time.Time  `json:"created"`
	Updated  time.Time  `json:"updated"`
}

func NewFileFrom(result files.FileResult) File {
	file := result.Meta
	chunks := make([]Document, len(file.Chunks))

	for i, fi := range file.Chunks {
		chunks[i] = Document{Chunk: int(fi.Chunk), Page: int(fi.Page)}
	}

	return File{
		Meta: FileMeta{
			Id:       uuid.UUID(file.Id).String(),
			Path:     string(file.Path),
			Provider: string(file.Provider),
			Created:  time.Time(file.Created),
			Updated:  time.Time(file.Updated),
			Chunks:   chunks,
		},
		Content: result.Content,
	}
}
