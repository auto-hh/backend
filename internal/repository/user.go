package repository

import (
	"context"

	"github.com/auto-hh/backend/internal/domain"
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

func (u *User) IsUserExistsByHHID(ctx context.Context, hhID uuid.UUID) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE hh_id = $1::UUID);`

	var exists bool

	executor := u.GetExecutor(ctx)

	err := executor.QueryRow(ctx, query, hhID).Scan(&exists)
	if err != nil {
		return false, domain.NewInternalServerError(
			domain.CodeInternalServerError,
			"database error",
			err,
		)
	}

	return exists, nil
}
