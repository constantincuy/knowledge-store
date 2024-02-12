package files

import (
	"context"
	"github.com/constantincuy/knowledgestore/internal/core/domain/file"
	"github.com/constantincuy/knowledgestore/internal/core/domain/knowledgebase"
	"github.com/constantincuy/knowledgestore/internal/core/service"
)

type SearchFilesReq struct {
	KnowledgeBase string
	Query         string
}

type SearchFilesRes struct {
	Files []file.File
}

func (s Service) Search(ctx context.Context, req SearchFilesReq) (SearchFilesRes, error) {
	knowledgeBase, err := knowledgebase.NewName(req.KnowledgeBase)
	if err != nil {
		return SearchFilesRes{}, service.NewError(service.ErrBadRequest, err)
	}

	data := make([]string, 1)
	data[0] = req.Query

	embed, err := s.embedding.Extract(data)
	if err != nil {
		return SearchFilesRes{}, service.NewError(service.ErrInternalFailure, err)
	}

	files, err := s.fileRepo.Search(ctx, knowledgeBase, embed.Vectors)
	if err != nil {
		return SearchFilesRes{}, service.NewError(service.ErrInternalFailure, err)
	}

	return SearchFilesRes{files}, nil
}
