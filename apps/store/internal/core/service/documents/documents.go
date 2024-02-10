package documents

import (
	"context"
	"github.com/constantincuy/knowledgestore/internal/ports"
)

type Api interface {
	Create(context.Context, AddDocumentReq) (AddDocumentRes, error)
}

type Service struct {
	docRepo ports.DocumentRepo
}

func NewService(docRepo ports.DocumentRepo) Service {
	return Service{docRepo}
}
