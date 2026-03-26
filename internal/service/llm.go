package service

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/auto-hh/backend/internal/domain"
	"github.com/auto-hh/backend/internal/repository"
	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

type LLM struct {
	repository repository.IProfile
	client     *http.Client
	llmPath    string
}

func NewLLM(repository repository.IProfile, client *http.Client) *LLM {
	return &LLM{repository, client}
}

func (llm *LLM) FindVacancies(ctx context.Context, userID uuid.UUID) error {

	data, err := llm.repository.GetProfileData(ctx, userID)
	if err != nil {
		return err
	}

	userInfo, err := json.Marshal(data)
	if err != nil {
		return domain.NewInternalServerError(domain.CodeInternalServerError, "Failed to marshal user info", err)
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, llm.llmPath+"/vacancies", bytes.NewReader(userInfo))
	if err != nil {
		return domain.NewInternalServerError(domain.CodeInternalServerError, "Failed to create request", err)
	}
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	response, err := llm.client.Do(request)
	if err != nil {
		return domain.NewInternalServerError(domain.CodeInternalServerError, "Failed to send request", err)
	}

	defer response.Body.Close()
	return nil
}
