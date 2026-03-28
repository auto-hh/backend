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
	_, err := middleware.GetUserID(ctx)
	if err != nil {
		return domain.MapAppError(ctx, err)
	}

	// vacancies, err := llm.service.FindVacancies(ctx.Request().Context(), userID)
	// if err != nil {
	// 	return domain.MapAppError(ctx, err)
	// }
	vacancies := []model.Vacancy{
        {
            JobTitle: "Frontend Developer",
            Salary: "150 000 - 220 000 ₽",
            City: "Москва",
            Body: "Разработка веб-приложений на React и TypeScript. Опыт работы от 2 лет.",
            Link: "https://hh.ru/vacancy/1",
            WorkFormat: "офис",
            Score: 1.0,
        },
        {
            JobTitle: "Backend Developer",
            Salary: "200 000 - 300 000 ₽",
            City: "Санкт-Петербург",
            Body: "Разработка микросервисов на Node.js и Python. Опыт работы от 3 лет.",
            Link: "https://hh.ru/vacancy/2",
            WorkFormat: "офис",
            Score: 0.85,
        },
        {
            JobTitle: "Fullstack Developer",
            Salary: "120 000 - 180 000 ₽",
            City: "Казань",
            Body: "Разработка fullstack-приложений для стартапов. React + Node.js.",
            Link: "https://hh.ru/vacancy/3",
            WorkFormat: "офис",
            Score: 1.0,
        },
        {
            JobTitle: "DevOps Engineer",
            Salary: "250 000 - 350 000 ₽",
            City: "Москва",
            Body: "Автоматизация инфраструктуры и CI/CD процессов. Kubernetes, Docker, AWS.",
            Link: "https://hh.ru/vacancy/4",
            WorkFormat: "офис",
            Score: 0.6,
        },
        {
            JobTitle: "QA Engineer",
            Salary: "100 000 - 160 000 ₽",
            City: "Новосибирск",
            Body: "Тестирование веб и мобильных приложений. Playwright, Jest, Postman.",
            Link: "https://hh.ru/vacancy/5",
            WorkFormat: "офис",
            Score: 1.0,
        },

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
	_, err := middleware.GetUserID(ctx)
	if err != nil {
		return domain.MapAppError(ctx, err)
	}

	// scores, err := llm.service.Analysis(ctx.Request().Context(), userID)
	// if err != nil {
	// 	return domain.MapAppError(ctx, err)
	// }
	scores := []model.Attribute{
		{ Text: "Senior", Weight: 1.0, IsWord: true },
		{ Text: " ", Weight: 0.0, IsWord: false },
		{ Text: "Python", Weight: 0.87, IsWord: true },
		{ Text: " ", Weight: 0.0, IsWord: false },
		{ Text: "Backend", Weight: 0.72, IsWord: true },
		{ Text: " ", Weight: 0.0, IsWord: false },
		{ Text: "Developer", Weight: 0.65, IsWord: true },
		{ Text: ".", Weight: 0.0, IsWord: false },
		{ Text: " ", Weight: 0.0, IsWord: false },
		{ Text: "Опыт", Weight: 0.15, IsWord: true },
		{ Text: ":", Weight: 0.0, IsWord: false },
		{ Text: " ", Weight: 0.0, IsWord: false },
		{ Text: "5", Weight: 0.08, IsWord: true },
		{ Text: " ", Weight: 0.0, IsWord: false },
		{ Text: "лет", Weight: 0.05, IsWord: true },
		{ Text: ".", Weight: 0.0, IsWord: false },
		{ Text: " ", Weight: 0.0, IsWord: false },
		{ Text: "Навыки", Weight: 0.12, IsWord: true },
		{ Text: ":", Weight: 0.0, IsWord: false },
		{ Text: " ", Weight: 0.0, IsWord: false },
		{ Text: "Python", Weight: 0.82, IsWord: true },
		{ Text: ",", Weight: 0.0, IsWord: false },
		{ Text: " ", Weight: 0.0, IsWord: false },
		{ Text: "FastAPI", Weight: 0.45, IsWord: true },
		{ Text: ",", Weight: 0.0, IsWord: false },
		{ Text: " ", Weight: 0.0, IsWord: false },
		{ Text: "PostgreSQL", Weight: 0.38, IsWord: true },
		{ Text: ",", Weight: 0.0, IsWord: false },
		{ Text: " ", Weight: 0.0, IsWord: false },
		{ Text: "Docker", Weight: 0.32, IsWord: true },
		{ Text: ",", Weight: 0.0, IsWord: false },
		{ Text: " ", Weight: 0.0, IsWord: false },
		{ Text: "Kafka", Weight: 0.28, IsWord: true },
		{ Text: ".", Weight: 0.0, IsWord: false },
		{ Text: " ", Weight: 0.0, IsWord: false },
		{ Text: "Образование", Weight: 0.1, IsWord: true },
		{ Text: ":", Weight: 0.0, IsWord: false },
		{ Text: " ", Weight: 0.0, IsWord: false },
		{ Text: "МГУ", Weight: 0.18, IsWord: true },
		{ Text: ",", Weight: 0.0, IsWord: false },
		{ Text: " ", Weight: 0.0, IsWord: false },
		{ Text: "2020", Weight: 0.05, IsWord: true },
		{ Text: ".", Weight: 0.0, IsWord: false },
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
