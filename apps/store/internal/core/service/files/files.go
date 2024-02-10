package files

import (
	"context"
	"github.com/constantincuy/knowledgestore/internal/ports"
)

type Api interface {
	GetFilesystem(context.Context, GetFilesystemReq) (GetFilesystemRes, error)
	Get(context.Context, GetFileReq) (GetFileRes, error)
}

type Service struct {
	fileRepo ports.FileRepo
}

func NewService(fileRepo ports.FileRepo) Service {
	return Service{fileRepo}
}
