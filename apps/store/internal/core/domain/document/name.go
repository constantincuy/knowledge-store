package document

import (
	"errors"
	"strings"
)

var (
	ErrEmptyDocumentName = errors.New("empty document name supplied")
)

type Name string

func NewName(un string) (Name, error) {
	un = strings.TrimSpace(un)
	if un == "" {
		return "", ErrEmptyDocumentName
	}

	return Name(un), nil
}
