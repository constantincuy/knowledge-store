package ports

import (
	"context"
	"github.com/constantincuy/knowledgestore/internal/core/domain/knowledgebase"
)

type KnowledgeBaseRepo interface {
	Add(ctx context.Context, u knowledgebase.KnowledgeBase) error
}
