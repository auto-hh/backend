package handler

import (
	"net/http"

	"github.com/auto-hh/backend/internal/domain"
	"github.com/auto-hh/backend/internal/middleware"
	"github.com/auto-hh/backend/internal/model"
	"github.com/auto-hh/backend/internal/service"
	"github.com/labstack/echo/v5"
)

type LLM struct {
	service service.ILLM
}

func NewLLM(service service.ILLM) *LLM {
	return &LLM{
		service: service,
	}
}

// FindVacancies
//
//	@Tags		llm
//	@Success	200	{object}	[]model.Vacancy
//	@Failure	401	{object}	domain.ErrorWrapper
//	@Failure	403	{object}	domain.ErrorWrapper
//	@Failure	500	{object}	domain.ErrorWrapper
//	@Security	BearerAuth
//	@Router		/llm/vacancies [post].
func (llm *LLM) FindVacancies(ctx *echo.Context) error {
	userID, err := middleware.GetUserID(ctx)
	if err != nil {
		return domain.MapAppError(ctx, err)
	}

	vacancies, err := llm.service.FindVacancies(ctx.Request().Context(), userID)
	if err != nil {
		return domain.MapAppError(ctx, err)
	}

	return domain.JSON(ctx, http.StatusOK, vacancies)
}

// Analysis
//
//	@Tags		llm
//	@Success	200	{object}	[]model.Attribute
//	@Failure	401	{object}	domain.ErrorWrapper
//	@Failure	403	{object}	domain.ErrorWrapper
//	@Failure	500	{object}	domain.ErrorWrapper
//	@Security	BearerAuth
//	@Router		/llm/analysis [post].
func (llm *LLM) Analysis(ctx *echo.Context) error {
	userID, err := middleware.GetUserID(ctx)
	if err != nil {
		return domain.MapAppError(ctx, err)
	}

	scores, err := llm.service.Analysis(ctx.Request().Context(), userID)
	if err != nil {
		return domain.MapAppError(ctx, err)
	}

	return domain.JSON(ctx, http.StatusOK, scores)
}

// GenerateCoverLetter
//
//	@Tags		llm
//	@Param		vacancy	body		model.Vacancy	true	"vacancy"
//	@Success	200		{object}	model.CoverLetter
//	@Failure	400		{object}	domain.ErrorWrapper
//	@Failure	401		{object}	domain.ErrorWrapper
//	@Failure	403		{object}	domain.ErrorWrapper
//	@Failure	500		{object}	domain.ErrorWrapper
//	@Security	BearerAuth
//	@Router		/llm/generate [post].
func (llm *LLM) GenerateCoverLetter(ctx *echo.Context) error {
	_, err := middleware.GetUserID(ctx)
	if err != nil {
		return domain.MapAppError(ctx, err)
	}

	var vacancy model.Vacancy
	err = ctx.Bind(&vacancy)
	if err != nil {
		err = domain.NewBadRequest(domain.CodeBadRequest, "unable to parse vacancy for generation", err)
		return domain.MapAppError(ctx, err)
	}

	// coverLetter, err := llm.service.GetCoverLetter(ctx.Request().Context(), userID, vacancy)
	// if err != nil {
	// 	return domain.MapAppError(ctx, err)
	// }

	coverLetter := model.CoverLetter{
		Letter: "THE BEST LETTER IN THE WORLD!!!",
		Status: "VERY GOOOD!!!!!!!!!!",
	}

	return domain.JSON(ctx, http.StatusOK, coverLetter)
}
