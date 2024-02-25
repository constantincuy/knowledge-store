package file

import "github.com/constantincuy/knowledgestore/internal/core/domain/common"

type File struct {
	Id       common.Id
	Path     Path
	Provider Provider
	Created  Created
	Updated  Updated
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
