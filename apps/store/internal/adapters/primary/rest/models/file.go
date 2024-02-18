package models

import (
	"github.com/constantincuy/knowledgestore/internal/core/domain/file"
	"github.com/google/uuid"
	"time"
)

type File struct {
	Id       string    `json:"id"`
	Path     string    `json:"path"`
	Provider string    `json:"provider"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}

func NewFileFrom(file file.File) File {
	return File{
		Id:       uuid.UUID(file.Id).String(),
		Path:     string(file.Path),
		Provider: string(file.Provider),
		Created:  time.Time(file.Created),
		Updated:  time.Time(file.Updated),
	}
}
