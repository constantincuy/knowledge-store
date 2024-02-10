package documents

import (
	"context"
	"github.com/constantincuy/knowledgestore/internal/core/domain/common"
	"github.com/constantincuy/knowledgestore/internal/core/domain/document"
	"github.com/constantincuy/knowledgestore/internal/core/domain/knowledgebase"
	"github.com/constantincuy/knowledgestore/internal/core/service"
	"github.com/google/uuid"
)

type AddDocumentReq struct {
	KnowledgeBase string
	FileId        uuid.UUID
	Chunk         int
	Embedding     []float32
}

type AddDocumentRes struct {
	Success bool
}

func (s Service) Create(ctx context.Context, req AddDocumentReq) (AddDocumentRes, error) {
	knowledgeBase, err := knowledgebase.NewName(req.KnowledgeBase)
	if err != nil {
		return AddDocumentRes{Success: false}, service.NewError(service.ErrBadRequest, err)
	}

	fileId, err := common.NewIdFrom(req.FileId)
	if err != nil {
		return AddDocumentRes{Success: false}, service.NewError(service.ErrBadRequest, err)
	}

	chunk, err := document.NewChunk(req.Chunk)
	if err != nil {
		return AddDocumentRes{Success: false}, service.NewError(service.ErrBadRequest, err)
	}

	embedding, err := document.NewEmbedding(req.Embedding)
	if err != nil {
		return AddDocumentRes{Success: false}, service.NewError(service.ErrBadRequest, err)
	}

	doc := document.NewDocument(fileId, chunk, embedding)

	if err = s.docRepo.Add(ctx, knowledgeBase, doc); err != nil {
		return AddDocumentRes{Success: false}, service.NewError(service.ErrInternalFailure, err)
	}

	return AddDocumentRes{Success: true}, nil
}
