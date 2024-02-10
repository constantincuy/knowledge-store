package document

import (
	"time"
)

type Created time.Time

func NewCreated() Created {
	return Created(time.Now())
}

func NewCreatedFrom(created time.Time) (Created, error) {
	return Created(created), nil
}
