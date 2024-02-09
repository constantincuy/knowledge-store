package file

import (
	"errors"
	"strings"
)

var (
	ErrEmptyPath = errors.New("empty path supplied")
)

type Path string

func NewPath(path string) (Path, error) {
	path = strings.TrimSpace(path)
	if path == "" {
		return "", ErrEmptyPath
	}

	return Path(path), nil
}
