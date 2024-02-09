package knowledgebase

import (
	"errors"
	"strings"
)

var (
	ErrEmptyKnowledgeBaseName   = errors.New("empty knowledge base name supplied")
	ErrInvalidKnowledgeBaseName = errors.New("invalid knowledge base name supplied only supported characters: [a-z] [A-Z] [0-9] _")
)

type Name string

func NewName(un string) (Name, error) {
	un = strings.TrimSpace(un)
	if un == "" {
		return "", ErrEmptyKnowledgeBaseName
	}

	return Name(un), nil
}
