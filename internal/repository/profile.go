package repository

import (
	"context"
	"errors"

	"github.com/auto-hh/backend/internal/domain"
	"github.com/auto-hh/backend/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

type Profile struct {
	*Executor
}

func NewProfile(e *Executor) *Profile {
	return &Profile{
		Executor: e,
	}
}

func (p *Profile) GetProfileData(ctx context.Context, userID uuid.UUID) (model.Profile, error) {
	query := `SELECT experience, job_title, grade, work_format, salary, city, about_me, recent_jobs FROM profiles WHERE id = $1::UUID`
	executor := p.GetExecutor(ctx)

	var data model.Profile

	err := executor.QueryRow(ctx, query, userID).Scan(
		&data.Experience,
		&data.JobTitle,
		&data.Grade,
		&data.WorkFormat,
		&data.Salary,
		&data.City,
		&data.AboutMe,
		&data.RecentJobs,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case pgerrcode.UniqueViolation:
				return model.Profile{}, domain.NewNotFound(domain.CodeNotFound, "not found user profile")
			}
		}

		return model.Profile{}, domain.NewInternalServerError(
			domain.CodeInternalServerError,
			"failed to get profile data",
			err,
		)
	}

	return data, nil
}

func (p *Profile) IsProfileExistsByUserID(ctx context.Context, userID uuid.UUID) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM profiles WHERE user_id = $1::UUID);`
	executor := p.GetExecutor(ctx)

	var exists bool
	err := executor.QueryRow(ctx, query, userID).Scan(
		&exists,
	)

	if err != nil {
		return false, domain.NewInternalServerError(domain.CodeInternalServerError, "failed to check if user exists", err)
	}

	return exists, nil
}
