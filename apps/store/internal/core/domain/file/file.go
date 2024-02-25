package file

import (
	"github.com/constantincuy/knowledgestore/internal/core/domain/common"
	"github.com/constantincuy/knowledgestore/internal/core/domain/document"
)

type File struct {
	Id       common.Id
	Path     Path
	Provider Provider
	Created  Created
	Updated  Updated
	Chunks   []document.Document
}

func NewFile(path Path, provider Provider, created Created, updated Updated) File {
	return File{
		Id:       common.NewId(),
		Path:     path,
		Provider: provider,
		Created:  created,
		Updated:  updated,
	}
}
