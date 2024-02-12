package knowledgebase

import (
	"errors"
	"strings"
)

var (
	ErrEmptyKnowledgeBaseName   = errors.New("empty knowledge base name supplied")
	ErrKnowledgeBaseNameTooLong = errors.New("knowledge base name is too long max 255 characters")
	ErrInvalidKnowledgeBaseName = errors.New("invalid knowledge base name supplied only supported characters: [a-z] [A-Z] [0-9] _")
)

type Name string

func NewName(un string) (Name, error) {
	un = strings.TrimSpace(un)
	if un == "" {
		return "", ErrEmptyKnowledgeBaseName
	}

	if len(un) > 255 {
		return "", ErrKnowledgeBaseNameTooLong
	}

	return Name(un), nil
}
