package repository

import (
	"context"
	"errors"

	"github.com/auto-hh/backend/internal/domain"
	"github.com/auto-hh/backend/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
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
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Profile{}, nil
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

func (p *Profile) InsertOrUpdate(ctx context.Context, userID uuid.UUID, profile model.Profile) error {
	query := `
		INSERT INTO profiles (user_id, experience, job_title, grade, work_format, salary, city, about_me, recent_jobs)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        ON CONFLICT (user_id)
        DO UPDATE SET
            experience = EXCLUDED.experience,
            job_title = EXCLUDED.job_title,
            grade = EXCLUDED.grade,
            work_format = EXCLUDED.work_format,
            salary = EXCLUDED.salary,
    		city = EXCLUDED.city,
            about_me = EXCLUDED.about_me,
            recent_jobs = EXCLUDED.recent_jobs;
	`

	_, err := p.GetExecutor(ctx).Exec(ctx, query, userID, profile.Experience, profile.JobTitle, profile.Grade, profile.WorkFormat, profile.Salary, profile.City, profile.AboutMe, profile.RecentJobs)

	if err != nil {
		return domain.NewInternalServerError(domain.CodeInternalServerError, "failed to insert or update user profile", err)
	}

	return nil
}
