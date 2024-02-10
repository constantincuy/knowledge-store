package ports

import (
	"context"
	"github.com/constantincuy/knowledgestore/internal/core/domain/file"
	"os"
)

type Storage interface {
	Provider() string
	GetChangedDocuments(ctx context.Context, filesystem file.Filesystem) (file.ChangeList, error)
	DownloadDocument(ctx context.Context, path file.Path, target *os.File)
}
