package ports

import (
	"context"
	"github.com/constantincuy/knowledgestore/internal/core/domain/document"
	"github.com/constantincuy/knowledgestore/internal/core/domain/knowledgebase"
)

type DocumentRepo interface {
	Add(ctx context.Context, knowledgeBase knowledgebase.Name, document document.Document) error
}
