package file

import (
	"time"
)

type Created time.Time

func NewCreated(created time.Time) (Created, error) {
	return Created(created), nil
}
