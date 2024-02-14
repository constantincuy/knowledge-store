package ports

import (
	"context"
	"github.com/constantincuy/knowledgestore/internal/core/domain/common"
	"github.com/constantincuy/knowledgestore/internal/core/domain/document"
	"github.com/constantincuy/knowledgestore/internal/core/domain/knowledgebase"
)

type DocumentRepo interface {
	Add(ctx context.Context, knowledgeBase knowledgebase.Name, document document.Document) error
	Delete(ctx context.Context, knowledgeBase knowledgebase.Name, fileId common.Id) error
}
