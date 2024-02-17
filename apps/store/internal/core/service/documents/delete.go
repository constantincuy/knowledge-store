package documents

import (
	"context"
	"github.com/constantincuy/knowledgestore/internal/core/domain/common"
	"github.com/constantincuy/knowledgestore/internal/core/domain/knowledgebase"
	"github.com/constantincuy/knowledgestore/internal/core/service"
	"github.com/google/uuid"
)

type DeleteDocumentReq struct {
	KnowledgeBase string
	FileId        uuid.UUID
}

type DeleteDocumentRes struct {
	Success bool
}

func (s Service) Delete(ctx context.Context, req DeleteDocumentReq) (DeleteDocumentRes, error) {
	knowledgeBase, err := knowledgebase.NewName(req.KnowledgeBase)
	if err != nil {
		return DeleteDocumentRes{Success: false}, service.NewError(service.ErrBadRequest, err)
	}

	fileId, err := common.NewIdFrom(req.FileId)
	if err != nil {
		return DeleteDocumentRes{Success: false}, service.NewError(service.ErrBadRequest, err)
	}

	if err = s.docRepo.Delete(ctx, knowledgeBase, fileId); err != nil {
		return DeleteDocumentRes{Success: false}, service.NewError(service.ErrInternalFailure, err)
	}

	return DeleteDocumentRes{Success: true}, nil
}
