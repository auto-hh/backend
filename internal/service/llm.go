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

func (llm *LLM) FindVacancies(ctx context.Context, userID uuid.UUID) ([]model.Vacancy, error) {

	requestLLM, err := llm.makeLLMRequest(ctx, userID, http.MethodPost, llm.llmPath+"/vacancies")
	if err != nil {
		return nil, err
	}

	response, err := llm.client.Do(requestLLM)
	if err != nil {
		return nil, domain.NewInternalServerError(domain.CodeInternalServerError, "Failed to send requestLLM", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, domain.NewInternalServerError(domain.CodeInternalServerError, "Failed to  receive response", err)
	}

	var respData []model.Vacancy
	err = json.NewDecoder(response.Body).Decode(&respData)
	if err != nil {
		return nil, domain.NewInternalServerError(domain.CodeInternalServerError, "Failed to read response body", err)
	}

	return respData, nil
}

func (llm *LLM) Analysis(ctx context.Context, userID uuid.UUID) ([]model.Attribute, error) {
	requestLLM, err := llm.makeLLMRequest(ctx, userID, http.MethodPost, llm.llmPath+"/analysis")
	if err != nil {
		return nil, err
	}
	response, err := llm.client.Do(requestLLM)
	if err != nil {
		return nil, domain.NewInternalServerError(domain.CodeInternalServerError, "Failed to send requestLLM", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, domain.NewInternalServerError(domain.CodeInternalServerError, "Failed to receive response", err)
	}

	var respData []model.Attribute
	err = json.NewDecoder(response.Body).Decode(&respData)
	if err != nil {
		return nil, domain.NewInternalServerError(domain.CodeInternalServerError, "Failed to read response body", err)
	}

	return respData, nil
}

func (llm *LLM) makeLLMRequest(ctx context.Context, userID uuid.UUID, method, customUrl string) (*http.Request, error) {
	rawUserInfo, err := llm.repository.GetProfileData(ctx, userID)
	if err != nil {
		return nil, err
	}

	userInfo, err := json.Marshal(rawUserInfo)
	if err != nil {
		return nil, domain.NewInternalServerError(domain.CodeInternalServerError, "Failed to marshal user info", err)
	}

	requestLLM, err := http.NewRequestWithContext(ctx, method, customUrl, bytes.NewReader(userInfo))
	if err != nil {
		return nil, domain.NewInternalServerError(domain.CodeInternalServerError, "Failed to create requestLLM", err)
	}

	requestLLM.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	return requestLLM, nil
}
