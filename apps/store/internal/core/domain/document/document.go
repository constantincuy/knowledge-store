package document

import (
	"github.com/constantincuy/knowledgestore/internal/core/domain/common"
)

type Document struct {
	Id        common.Id
	FileId    common.Id
	Page      PageNumber
	Chunk     Chunk
	Created   Created
	Embedding Embedding
}

func NewDocument(fileId common.Id, chunk Chunk, embedding Embedding) Document {
	number, _ := NewPageNumber(1)
	return Document{
		Id:        common.NewId(),
		FileId:    fileId,
		Page:      number,
		Chunk:     chunk,
		Created:   NewCreated(),
		Embedding: embedding,
	}
}
