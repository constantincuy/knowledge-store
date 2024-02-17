package postgres

import (
	"context"
	"github.com/constantincuy/knowledgestore/internal/core/domain/common"
	"github.com/constantincuy/knowledgestore/internal/core/domain/document"
	"github.com/constantincuy/knowledgestore/internal/core/domain/knowledgebase"
	"github.com/constantincuy/knowledgestore/internal/ports"
	"github.com/google/uuid"
	"strconv"
	"strings"
	"time"
)

type DocumentRepo struct {
	provider ports.DatabaseProvider
}

func (d DocumentRepo) Delete(ctx context.Context, knowledgeBase knowledgebase.Name, fileId common.Id) error {
	db, err := d.provider.GetDatabase(string(knowledgeBase))
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.PrepareContext(ctx, "DELETE FROM document_collection WHERE file_id = (SELECT id FROM filesystem WHERE unique_id = $1)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, uuid.UUID(fileId).String())

	return err
}

func float32ArrayToString(arr []float32) string {
	var strValues []string

	for _, v := range arr {
		strValues = append(strValues, strconv.FormatFloat(float64(v), 'f', -1, 32))
	}

	return "[" + strings.Join(strValues, ", ") + "]"
}

func (d DocumentRepo) Add(ctx context.Context, knowledgeBase knowledgebase.Name, doc document.Document) error {
	db, err := d.provider.GetDatabase(string(knowledgeBase))
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.PrepareContext(ctx, "INSERT INTO document_collection(unique_id, file_id, page, chunk, created, embedding) VALUES ($1, (SELECT Id FROM filesystem WHERE unique_id = $2), $3, $4, $5, $6)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, uuid.UUID(doc.Id).String(), uuid.UUID(doc.FileId).String(), doc.Page, doc.Chunk, time.Time(doc.Created), float32ArrayToString(doc.Embedding))

	return err
}

func NewDocumentRepo(provider ports.DatabaseProvider) (DocumentRepo, error) {
	return DocumentRepo{provider}, nil
}
