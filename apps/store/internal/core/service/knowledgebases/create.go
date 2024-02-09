package knowledgebases

import (
	"context"
	"fmt"
	"github.com/constantincuy/knowledgestore/internal/core/domain/knowledgebase"
	"github.com/constantincuy/knowledgestore/internal/core/service"
)

type CreateKnowledgeBaseReq struct {
	Name string
}

type CreateKnowledgeBaseResp struct {
	Success bool
}

func (s Service) Create(ctx context.Context, req CreateKnowledgeBaseReq) (CreateKnowledgeBaseResp, error) {
	name, err := knowledgebase.NewName(req.Name)
	if err != nil {
		return CreateKnowledgeBaseResp{Success: false}, service.NewError(service.ErrBadRequest, err)
	}

	base := knowledgebase.New(name)

	if err := s.baseRepo.Add(ctx, base); err != nil {
		return CreateKnowledgeBaseResp{Success: false}, service.NewError(service.ErrBadRequest, fmt.Errorf("failed to add a knowledge base: %w", err))
	}

	return CreateKnowledgeBaseResp{Success: true}, nil
}
