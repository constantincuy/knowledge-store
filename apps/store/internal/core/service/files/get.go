package files

import (
	"context"
	"github.com/constantincuy/knowledgestore/internal/core/domain/file"
)

type GetFileReq struct {
	KnowledgeBase string
	Provider      string
	Path          string
}

type GetFileRes struct {
	File file.File
}

func (s Service) Get(ctx context.Context, req GetFileReq) (GetFileRes, error) {
	return GetFileRes{}, nil
}
