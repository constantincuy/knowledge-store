package knowledgebases

import (
	"context"
	"github.com/constantincuy/knowledgestore/internal/ports"
)

type Api interface {
	Create(context.Context, CreateKnowledgeBaseReq) (CreateKnowledgeBaseResp, error)
}

type Service struct {
	baseRepo ports.KnowledgeBaseRepo
}

func NewService(baseRepo ports.KnowledgeBaseRepo) Service {
	return Service{baseRepo}
}
