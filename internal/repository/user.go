package repository

import (
	"context"

	"github.com/google/uuid"
)

type User struct {
	*Executor
}

func NewUser(executor *Executor) *User {
	return &User{
		Executor: executor,
	}
}

func IsUserExistsByHHID(ctx context.Context, hhID uuid.UUID) (bool, error) {
	return false, nil
}
