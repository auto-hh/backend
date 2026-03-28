package service

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/auto-hh/backend/internal/domain"
	"github.com/auto-hh/backend/internal/model"
	"github.com/auto-hh/backend/internal/repository"
	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

type LLM struct {
	repository repository.IProfile
	client     *http.Client
	llmPath    string
}

func NewLLM(repository repository.IProfile, client *http.Client, llmPath string) *LLM {
	return &LLM{repository, client, llmPath}
}

func (llm *LLM) FindVacancies(ctx context.Context, userID uuid.UUID) ([]model.Vacancy, error) {
	rawUserInfo, err := llm.repository.GetProfileData(ctx, userID)
	if err != nil {
		return nil, err
	}

	requestLLM, err := llm.makeLLMRequest(ctx, userID, http.MethodPost, llm.llmPath+"/search", rawUserInfo)
	if err != nil {
		return nil, err
	}

	response, err := llm.client.Do(requestLLM)
	if err != nil {
		return nil, domain.NewInternalServerError(
			domain.CodeInternalServerError,
			"Failed to send requestLLM",
			err,
		)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		responseData, _ := io.ReadAll(response.Body)
		return nil, domain.NewInternalServerError(
			domain.CodeInternalServerError,
			"vacancies response status is not ok: "+string(responseData),
			err,
		)
	}

	var respData []model.Vacancy

	err = json.NewDecoder(response.Body).Decode(&respData)
	if err != nil {
		return nil, domain.NewInternalServerError(
			domain.CodeInternalServerError,
			"Failed to read response body",
			err,
		)
	}

	return respData, nil
}

func (llm *LLM) Analysis(ctx context.Context, userID uuid.UUID) ([]model.Attribute, error) {
	rawUserInfo, err := llm.repository.GetProfileData(ctx, userID)
	if err != nil {
		return nil, err
	}

	requestLLM, err := llm.makeLLMRequest(ctx, userID, http.MethodPost, llm.llmPath+"/analyze", rawUserInfo)
	if err != nil {
		return nil, err
	}

	response, err := llm.client.Do(requestLLM)
	if err != nil {
		return nil, domain.NewInternalServerError(
			domain.CodeInternalServerError,
			"Failed to send requestLLM",
			err,
		)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		responseData, _ := io.ReadAll(response.Body)
		return nil, domain.NewInternalServerError(
			domain.CodeInternalServerError,
			"analyze response status is not ok: "+string(responseData),
			err,
		)
	}

	var respData []model.Attribute

	err = json.NewDecoder(response.Body).Decode(&respData)
	if err != nil {
		responseData, _ := io.ReadAll(response.Body)
		return nil, domain.NewInternalServerError(
			domain.CodeInternalServerError,
			"unable to decode analyze data: "+string(responseData),
			err,
		)
	}

	return respData, nil
}

func (llm *LLM) GetCoverLetter(ctx context.Context, userID uuid.UUID, vacancy model.Vacancy) (model.CoverLetter, error) {
	rawUserInfo, err := llm.repository.GetProfileData(ctx, userID)
	if err != nil {
		return model.CoverLetter{}, err
	}
	generateRequest := model.GenerateRequest{
		Resume: rawUserInfo,
		Vacancy: vacancy,
	}
	requestLLM, err := llm.makeLLMRequest(ctx, userID, http.MethodPost, llm.llmPath+"/generate", generateRequest)
	if err != nil {
		return model.CoverLetter{}, err
	}

	response, err := llm.client.Do(requestLLM)
	if err != nil {
		return model.CoverLetter{}, domain.NewInternalServerError(
			domain.CodeInternalServerError,
			"Failed to send requestLLM",
			err,
		)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		responseData, _ := io.ReadAll(response.Body)
		return model.CoverLetter{}, domain.NewInternalServerError(
			domain.CodeInternalServerError,
			"generate response status is not ok: "+string(responseData),
			err,
		)
	}

	var coverLetter model.CoverLetter

	err = json.NewDecoder(response.Body).Decode(&coverLetter)
	if err != nil {
		return model.CoverLetter{}, domain.NewInternalServerError(
			domain.CodeInternalServerError,
			"Failed to decode cover letter",
			err,
		)
	}

	return coverLetter, nil
}

func (llm *LLM) makeLLMRequest(
	ctx context.Context,
	userID uuid.UUID,
	method, customURL string,
	data any,
) (*http.Request, error) {
	requestBody, err := json.Marshal(data)
	if err != nil {
		return nil, domain.NewInternalServerError(
			domain.CodeInternalServerError,
			"Failed to marshal llm request body",
			err,
		)
	}

	requestLLM, err := http.NewRequestWithContext(ctx, method, customURL, bytes.NewReader(requestBody))
	if err != nil {
		return nil, domain.NewInternalServerError(
			domain.CodeInternalServerError,
			"Failed to create requestLLM",
			err,
		)
	}

	requestLLM.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	return requestLLM, nil
}
