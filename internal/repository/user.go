package repository

import (
	"context"

	"github.com/auto-hh/backend/internal/domain"
	"github.com/auto-hh/backend/internal/model"
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

func (u *User) GetOrCreate(ctx context.Context, userData *model.UserData) (uuid.UUID, error) {
	query := `
		INSERT INTO users (hh_id, first_name, last_name)
		VALUES ($1::VARCHAR(32), $2::VARCHAR(32), $3::VARCHAR(32))
		ON CONFLICT (hh_id) DO UPDATE SET first_name = users.first_name, last_name = users.last_name
		RETURNING id;
	`

	var userID uuid.UUID

	err := u.GetExecutor(ctx).QueryRow(ctx, query, userData.ID, userData.FirstName, userData.LastName).Scan(&userID)
	if err != nil {
		return uuid.UUID{}, domain.NewInternalServerError(
			domain.CodeInternalServerError,
			"database error",
			err,
		)
	}

	return userID, nil
}
