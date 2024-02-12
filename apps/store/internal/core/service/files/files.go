package files

import (
	"context"
	"github.com/constantincuy/knowledgestore/internal/ports"
)

type Api interface {
	GetFilesystem(context.Context, GetFilesystemReq) (GetFilesystemRes, error)
	Get(context.Context, GetFileReq) (GetFileRes, error)
	Search(context.Context, SearchFilesReq) (SearchFilesRes, error)
}

type Service struct {
	fileRepo  ports.FileRepo
	embedding ports.EmbeddingExtractor
}

func NewService(fileRepo ports.FileRepo, embedding ports.EmbeddingExtractor) Service {
	return Service{fileRepo, embedding}
}
