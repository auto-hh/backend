package domain

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v5"
)

type ErrorResponse struct {
	code AppErrorCode
}

func JSON(ctx *echo.Context, code int, data any) error {
	return ctx.JSON(code, data)
}

func MapAppError(ctx *echo.Context, err error) error {
	mapAppErr := map[AppErrorType]int{
		TypeBadRequest:          http.StatusBadRequest,
		TypeUnauthorized:        http.StatusUnauthorized,
		TypeForbidden:           http.StatusForbidden,
		TypeInternalServerError: http.StatusInternalServerError,
	}

	var appError AppError
	if errors.As(err, &appError) {
		status, ok := mapAppErr[appError.errorType]
		if ok {
			return JSON(ctx, status, map[string]ErrorResponse{"error": {code: appError.code}})
		}

		return JSON(
			ctx,
			http.StatusInternalServerError,
			map[string]ErrorResponse{"error": {code: CodeInternalServerError}},
		)
	}

	return JSON(
		ctx,
		http.StatusInternalServerError,
		map[string]ErrorResponse{"error": {code: CodeInternalServerError}},
	)
}
