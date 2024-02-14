package ports

import (
	"context"
	"github.com/constantincuy/knowledgestore/internal/core/domain/document"
	"github.com/constantincuy/knowledgestore/internal/core/domain/file"
	"github.com/constantincuy/knowledgestore/internal/core/domain/knowledgebase"
)

type FileRepo interface {
	Add(ctx context.Context, knowledgeBase knowledgebase.Name, file file.File) error
	Delete(ctx context.Context, knowledgeBase knowledgebase.Name, provider file.Provider, path file.Path) error
	Get(ctx context.Context, knowledgeBase knowledgebase.Name, path file.Path) (file.File, error)
	Search(ctx context.Context, knowledgeBase knowledgebase.Name, embedding document.Embedding) ([]file.File, error)
	GetFromProvider(ctx context.Context, knowledgeBase knowledgebase.Name, provider file.Provider, path file.Path) (file.File, error)
	GetAllProviderFiles(ctx context.Context, knowledgeBase knowledgebase.Name, provider file.Provider) ([]file.File, error)
}
