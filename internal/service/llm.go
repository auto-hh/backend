package service

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/auto-hh/backend/internal/domain"
	"github.com/auto-hh/backend/internal/model"
	"github.com/auto-hh/backend/internal/repository"
	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

type LLM struct {
	repository   repository.IProfile
	client       *http.Client
	llmPath      string
	frontendPath string
}

func NewLLM(repository repository.IProfile, client *http.Client, llmPath, frontendPath string) *LLM {
	return &LLM{repository, client, llmPath, frontendPath}
}

func (llm *LLM) FindVacancies(ctx context.Context, userID uuid.UUID) error {

	rawUserInfo, err := llm.repository.GetProfileData(ctx, userID)
	if err != nil {
		return err
	}

	userInfo, err := json.Marshal(rawUserInfo)
	if err != nil {
		return domain.NewInternalServerError(domain.CodeInternalServerError, "Failed to marshal user info", err)
	}

	requestLLM, err := http.NewRequestWithContext(ctx, http.MethodGet, llm.llmPath+"/vacancies", bytes.NewReader(userInfo))
	if err != nil {
		return domain.NewInternalServerError(domain.CodeInternalServerError, "Failed to create requestLLM", err)
	}
	requestLLM.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	response, err := llm.client.Do(requestLLM)
	if err != nil {
		return domain.NewInternalServerError(domain.CodeInternalServerError, "Failed to send requestLLM", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return domain.NewInternalServerError(domain.CodeInternalServerError, "Failed to  receive response", err)
	}

	var respData []model.Vacancy
	err = json.NewDecoder(response.Body).Decode(&respData)
	if err != nil {
		return domain.NewInternalServerError(domain.CodeInternalServerError, "Failed to read response body", err)
	}

	jsonVacancies, err := json.Marshal(respData)
	if err != nil {
		return domain.NewInternalServerError(domain.CodeInternalServerError, "Failed to read make json", err)
	}

	requestFrontend, err := http.NewRequestWithContext(ctx, http.MethodPost, llm.frontendPath+"/vacancies", bytes.NewReader(jsonVacancies))
	if err != nil {
		return domain.NewInternalServerError(domain.CodeInternalServerError, "Failed to create request to frontend", err)
	}

	requestFrontend.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	responseFrontend, err := llm.client.Do(requestFrontend)
	if err != nil {
		return domain.NewInternalServerError(domain.CodeInternalServerError, "Failed to send Vacancies to frontend", err)
	}

	if responseFrontend.StatusCode != http.StatusOK {
		return domain.NewInternalServerError(domain.CodeInternalServerError, "Failed to  receive frontend response", err)
	}

	return nil
}
