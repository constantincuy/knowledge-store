package file

import (
	"time"
)

type Updated time.Time

func NewUpdated(updated time.Time) (Updated, error) {
	return Updated(updated), nil
}
