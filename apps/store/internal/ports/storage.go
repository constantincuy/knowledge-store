package ports

import "context"

type Storage interface {
	Provider() string
	GetChangedDocuments(ctx context.Context) ([]string, error)
	GetDocument(ctx context.Context, path string)
}
