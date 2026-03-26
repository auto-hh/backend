package repository

import (
	"context"

	"github.com/auto-hh/backend/internal/domain"
	"github.com/auto-hh/backend/internal/model"
	"github.com/google/uuid"
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
		return model.Profile{}, domain.NewInternalServerError(domain.CodeInternalServerError, "failed to get profile data", err)
	}

	return data, nil
}
