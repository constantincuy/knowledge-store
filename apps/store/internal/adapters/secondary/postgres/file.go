package postgres

import (
	"context"
	"github.com/constantincuy/knowledgestore/internal/core/domain/common"
	"github.com/constantincuy/knowledgestore/internal/core/domain/document"
	"github.com/constantincuy/knowledgestore/internal/core/domain/file"
	"github.com/constantincuy/knowledgestore/internal/core/domain/knowledgebase"
	"github.com/constantincuy/knowledgestore/internal/ports"
	"github.com/google/uuid"
	"time"
)

type FileRepo struct {
	provider ports.DatabaseProvider
}

func (f FileRepo) Add(ctx context.Context, knowledgeBase knowledgebase.Name, fi file.File) error {
	db, err := f.provider.GetDatabase(string(knowledgeBase))
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.PrepareContext(ctx, "INSERT INTO filesystem(unique_id, provider, path, created, updated) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, uuid.UUID(fi.Id).String(), fi.Provider, fi.Path, time.Time(fi.Created), time.Time(fi.Updated))

	return err
}
func (f FileRepo) Get(ctx context.Context, knowledgeBase knowledgebase.Name, path file.Path) (file.File, error) {
	db, err := f.provider.GetDatabase(string(knowledgeBase))
	if err != nil {
		return file.File{}, err
	}
	defer db.Close()

	stmt, err := db.PrepareContext(ctx, "SELECT provider, path, created, updated FROM filesystem WHERE path = $2;")
	if err != nil {
		return file.File{}, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, string(path))
	if err != nil {
		return file.File{}, err
	}
	defer rows.Close()

	if rows.Next() {
		f := file.File{}
		rows.Scan(&f.Provider, &f.Path, &f.Created, &f.Updated)
		return f, nil
	}

	return file.File{}, nil
}

func (f FileRepo) Search(ctx context.Context, knowledgeBase knowledgebase.Name, embedding document.Embedding) ([]file.File, error) {
	result := make([]file.File, 0)
	db, err := f.provider.GetDatabase(string(knowledgeBase))
	if err != nil {
		return result, err
	}
	defer db.Close()

	stmt, err := db.PrepareContext(ctx, "SELECT f.unique_id, f.path, f.provider FROM document_collection dc JOIN filesystem f ON f.id = dc.file_id WHERE 1 - (embedding <=> $1) >= 0.5")
	if err != nil {
		return result, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, float32ArrayToString(embedding))
	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		fi := file.File{}
		var uid string
		err := rows.Scan(&uid, &fi.Path, &fi.Provider)
		if err != nil {
			return result, err
		}
		id, err := uuid.Parse(uid)
		if err == nil {
			fi.Id, err = common.NewIdFrom(id)
			result = append(result, fi)
		}
	}

	return result, err
}

func (f FileRepo) GetFromProvider(ctx context.Context, knowledgeBase knowledgebase.Name, provider file.Provider, path file.Path) (file.File, error) {
	db, err := f.provider.GetDatabase(string(knowledgeBase))
	if err != nil {
		return file.File{}, err
	}
	defer db.Close()

	stmt, err := db.PrepareContext(ctx, "SELECT provider, path, created, updated FROM filesystem WHERE provider = $1 AND path = $2;")
	if err != nil {
		return file.File{}, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, string(provider), string(path))
	if err != nil {
		return file.File{}, err
	}
	defer rows.Close()

	if rows.Next() {
		f := file.File{}
		rows.Scan(&f.Provider, &f.Path, &f.Created, &f.Updated)
		return f, nil
	}

	return file.File{}, nil
}

func (f FileRepo) GetAllProviderFiles(ctx context.Context, knowledgeBase knowledgebase.Name, provider file.Provider) ([]file.File, error) {
	files := make([]file.File, 0)
	db, err := f.provider.GetDatabase(string(knowledgeBase))
	if err != nil {
		return files, err
	}
	defer db.Close()

	stmt, err := db.PrepareContext(ctx, "SELECT provider, path, created, updated FROM filesystem WHERE provider = $1;")
	if err != nil {
		return files, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, string(provider))
	if err != nil {
		return files, err
	}
	defer rows.Close()

	for rows.Next() {
		f := file.File{}
		rows.Scan(&f.Provider, &f.Path, &f.Created, &f.Updated)
		files = append(files, f)
	}

	return files, nil
}

func NewFileRepo(provider ports.DatabaseProvider) (FileRepo, error) {
	return FileRepo{provider}, nil
}
