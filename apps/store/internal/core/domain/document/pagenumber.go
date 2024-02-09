package document

import (
	"errors"
)

var (
	ErrInvalidPageNumber = errors.New("invalid page number supplied")
)

type PageNumber int

func NewPageNumber(number int) (PageNumber, error) {
	if number < 1 {
		return 1, ErrInvalidPageNumber
	}

	return PageNumber(number), nil
}
