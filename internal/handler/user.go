package handler

import (
	"net/http"

	"github.com/auto-hh/backend/internal/domain"
	"github.com/auto-hh/backend/internal/middleware"
	"github.com/auto-hh/backend/internal/service"
	"github.com/labstack/echo/v5"
)

type User struct {
	service service.IUser
}

func NewUser(service service.IUser) *User {
	return &User{
		service: service,
	}
}

// @Tags         user
// @Success      204
// @Failure      401 {object} domain.ErrorWrapper
// @Failure      403 {object} domain.ErrorWrapper
// @Security      BearerAuth
// @Router       /user/me [get]
func (u *User) Me(ctx *echo.Context) error {
	return ctx.NoContent(http.StatusNoContent)
}

// @Tags         user
// @Success      204
// @Failure      401 {object} domain.ErrorWrapper
// @Failure      403 {object} domain.ErrorWrapper
// @Failure      404 {object} domain.ErrorWrapper
// @Failure      500 {object} domain.ErrorWrapper
// @Security      BearerAuth
// @Router       /user/has-profile [get]
func (u *User) HasProfile(ctx *echo.Context) error {

	userID, err := middleware.GetUserID(ctx)
	if err != nil {
		return domain.MapAppError(ctx, err)
	}
	exists, err := u.service.IsProfileExistsByUserID(ctx.Request().Context(), userID)

	if err != nil {
		return domain.MapAppError(ctx, err)
	}

	if !exists {
		return ctx.NoContent(http.StatusNotFound)
	}

	return ctx.NoContent(http.StatusNoContent)
}

// @Tags         user
// @Success      200 {object} model.Profile
// @Failure      401 {object} domain.ErrorWrapper
// @Failure      403 {object} domain.ErrorWrapper
// @Failure      404 {object} domain.ErrorWrapper
// @Failure      500 {object} domain.ErrorWrapper
// @Security      BearerAuth
// @Router       /user/profile [get]
func (u *User) Profile(ctx *echo.Context) error {

	userID, err := middleware.GetUserID(ctx)
	if err != nil {
		return domain.MapAppError(ctx, err)
	}

	profileInfo, err := u.service.GetUserInfo(ctx.Request().Context(), userID)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, profileInfo)
}

// @Tags         user
// @Success      200 {object}
// @Failure      400 {object} domain.ErrorWrapper
// @Failure      401 {object} domain.ErrorWrapper
// @Failure      403 {object} domain.ErrorWrapper
// @Failure      500 {object} domain.ErrorWrapper
// @Security      BearerAuth
// @Router       /user/profile [post]
// func (u *User) UpdateProfile(ctx *echo.Context) error {
// 	userID, err := middleware.GetUserID(ctx)
// 	if err != nil {
// 		return domain.MapAppError(ctx, err)
// 	}

// 	var profile model.Profile
// 	err = ctx.Bind(&profile)
// 	if err != nil {
// 		err = domain.NewBadRequest(domain.CodeBadRequest, "can not parse update profile data", err)
// 		return domain.MapAppError(ctx, err)
// 	}

// 	err = u.service.UpdateUserInfo(ctx.Request().Context(), userID)
// 	if err != nil {
// 		return domain.MapAppError(ctx, err)
// 	}

// 	return ctx.NoContent(http.StatusOK)
// }
