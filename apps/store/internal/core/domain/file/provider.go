package file

import (
	"errors"
	"strings"
)

var (
	ErrEmptyProvider = errors.New("empty provider supplied")
)

type Provider string

func NewProvider(provider string) (Provider, error) {
	provider = strings.TrimSpace(provider)
	if provider == "" {
		return "", ErrEmptyProvider
	}

	return Provider(provider), nil
}
