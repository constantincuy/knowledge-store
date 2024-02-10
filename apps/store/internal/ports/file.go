package ports

import (
	"context"
	"github.com/constantincuy/knowledgestore/internal/core/domain/file"
	"github.com/constantincuy/knowledgestore/internal/core/domain/knowledgebase"
)

type FileRepo interface {
	Add(ctx context.Context, knowledgeBase knowledgebase.Name, file file.File) error
	Get(ctx context.Context, knowledgeBase knowledgebase.Name, path file.Path) (file.File, error)
	GetFromProvider(ctx context.Context, knowledgeBase knowledgebase.Name, provider file.Provider, path file.Path) (file.File, error)
	GetAllProviderFiles(ctx context.Context, knowledgeBase knowledgebase.Name, provider file.Provider) ([]file.File, error)
}
