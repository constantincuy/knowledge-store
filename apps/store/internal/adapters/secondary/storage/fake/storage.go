package fake

import "context"

type Storage struct {
}

func (s Storage) Provider() string {
	return "fake_provider_1"
}

func (s Storage) GetChangedDocuments(ctx context.Context) ([]string, error) {
	return []string{"my/path/1.txt", "my/path/2.txt", "my/path/3.txt"}, nil
}

func (s Storage) GetDocument(ctx context.Context, path string) {
	//TODO implement me
	panic("implement me")
}
