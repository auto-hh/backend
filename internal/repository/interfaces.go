package repository

import (
	"context"

	"github.com/auto-hh/backend/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type repositoryCtxtKey string

const KeyTx repositoryCtxtKey = "pgx_tx"

type IExecutor interface {
	Exec(
		ctx context.Context,
		sql string,
		arguments ...any,
	) (commandTag pgconn.CommandTag, err error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

type TransactionManager interface {
	WithTransaction(ctx context.Context, function func(ctx context.Context) error) error
}

type IUser interface {
	IsUserExistsByHHID(ctx context.Context, hhID uuid.UUID) (bool, error)
	GetOrCreate(ctx context.Context, userData *model.UserData) (uuid.UUID, error)
}

type IProfile interface {
	GetProfileData(ctx context.Context, userID uuid.UUID) (model.Profile, error)
}
