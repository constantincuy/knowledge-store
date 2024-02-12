package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/constantincuy/knowledgestore/internal/core/domain/knowledgebase"
	"github.com/constantincuy/knowledgestore/internal/ports"
)

var (
	ErrAlreadyExists = errors.New("duplicate entry")
)

type KnowledgeBaseRepo struct {
	provider ports.DatabaseProvider
}

func (kb KnowledgeBaseRepo) GetAll(ctx context.Context) ([]string, error) {
	result := make([]string, 0)
	kbs, err := kb.provider.GetDatabase("kb_system")
	if err != nil {
		return result, err
	}
	defer kbs.Close()

	rows, err := kbs.QueryContext(ctx, "SELECT * FROM knowledge_bases")
	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		var val string
		rows.Scan(&val)
		result = append(result, val)
	}

	return result, nil
}

func NewKnowledgeBaseRepo(provider ports.DatabaseProvider) (KnowledgeBaseRepo, error) {
	return KnowledgeBaseRepo{provider}, nil
}

func (kb KnowledgeBaseRepo) Add(ctx context.Context, u knowledgebase.KnowledgeBase) error {
	global, err := kb.provider.GetDatabase("")

	if err != nil {
		if global != nil {
			global.Close()
		}

		return err
	}
	defer global.Close()

	stmt, err := global.PrepareContext(ctx, "SELECT 1 as counter FROM pg_database WHERE datname = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, string(u.Name))
	if err != nil {
		return err
	}
	defer rows.Close()

	ex := rows.Next()
	count := 0
	if ex {
		if err = rows.Scan(&count); err != nil {
			return err
		}
	}

	if count == 1 {
		return fmt.Errorf("%s: %s", ErrAlreadyExists, errors.New(fmt.Sprintf("%s already exists", string(u.Name))))
	}

	_, err = global.ExecContext(ctx, "CREATE DATABASE "+string(u.Name))

	if err != nil {
		return err
	}

	newDB, err := kb.provider.GetDatabase(string(u.Name))

	if err != nil {
		return err
	}
	defer newDB.Close()

	_, err = newDB.ExecContext(ctx, "CREATE EXTENSION vector")

	if err != nil {
		return err
	}

	_, err = newDB.ExecContext(ctx, "CREATE TABLE filesystem (id BIGSERIAL PRIMARY KEY, unique_id varchar(42), provider VARCHAR(255), path VARCHAR(1024), created TIMESTAMP, updated TIMESTAMP)")

	if err != nil {
		return err
	}

	_, err = newDB.ExecContext(ctx, `
CREATE TABLE document_collection (
    id BIGSERIAL PRIMARY KEY,
    unique_id varchar(42),
    file_id BIGSERIAL,
    page INT,
    chunk INT,
    created TIMESTAMP,
    embedding VECTOR,
    CONSTRAINT fk_document_file
      FOREIGN KEY(file_id)
	    REFERENCES filesystem(id)
)
`)

	if err != nil {
		return err
	}

	sys, err := kb.provider.GetDatabase("kb_system")
	if err != nil {
		return err
	}
	defer sys.Close()

	stmt, err = sys.PrepareContext(ctx, "INSERT INTO knowledge_bases VALUES ($1)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, string(u.Name))
	if err != nil {
		return err
	}

	return nil
}
