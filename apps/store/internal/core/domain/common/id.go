package common

import (
	"github.com/google/uuid"
)

type Id uuid.UUID

func NewId() Id {
	return Id(uuid.New())
}

func NewIdFrom(uid uuid.UUID) (Id, error) {
	return Id(uid), nil
}
