package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/auto-hh/backend/internal/domain"
	"github.com/auto-hh/backend/internal/model"
	"github.com/auto-hh/backend/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v5"
)

type Auth struct {
	txManager repository.TransactionManager
	repositoryUser repository.IUser
	client *http.Client
	secretKey []byte
	clientID string
	clientSecret string
	redirectURI string
	appName string
	appVersion string
	devContact string
}

func NewAuth(repositoryUser repository.IUser, client *http.Client, secretKey []byte, clientID, clientSecret, redirectURI, appName, appVersion, devContact string) *Auth {
	return &Auth{
		repositoryUser: repositoryUser,
		client: client,
		secretKey: secretKey,
		clientID: clientID,
		redirectURI: redirectURI,
		appName: appName,
		appVersion: appVersion,
		devContact: devContact,
	}
}

func (a *Auth) Begin() (string, *url.URL, error) {
	state, err := generateState()
	if err != nil {
		return "", nil, domain.NewInternalServerError(domain.CodeInternalServerError, "error in state generation", err)
	}

	stateData := model.JWTStateData{
		State: state,
	}
	stateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, stateData)
	stateTokenSigned, err := stateToken.SignedString(a.secretKey)
	if err != nil {
		return "", nil, domain.NewInternalServerError(domain.CodeInternalServerError, "error in state jwt token signing", err)
	}

	query := url.Values{
		"response_type": {"code"},
		"client_id": {a.clientID},
		"redirect_uri": {a.redirectURI},
		"state": {state},
	}

	redirectURL := &url.URL{
		Scheme: "https",
		Host: "hh.ru",
		Path: "/oauth/authorize",
		RawQuery: query.Encode(),
	}

	return stateTokenSigned, redirectURL, nil
}

func (a *Auth) Complete(ctx context.Context, stateJWTToken string, complete model.Complete) (authJWTToken string, err error) {
	var stateData model.JWTStateData
	token, err := jwt.ParseWithClaims(stateJWTToken, &stateData, func(t *jwt.Token) (any, error) {return a.secretKey, nil})
	if err != nil {
		err = domain.NewBadRequest(domain.CodeBadRequest, "state jwt token decode error", err)
		return
	}
	if !token.Valid {
		err = domain.NewBadRequest(domain.CodeBadRequest, "state jwt token invalid", err)
		return
	}
	if stateData.State != complete.State {
		err = domain.NewBadRequest(domain.CodeBadRequest, "states are not the same")
		return
	}
	
	hhData, err := a.getHHData(ctx, complete.Code)
	if err != nil {
		return
	}

	userData, err := a.getUserData(ctx, hhData.AccessToken)
	if err != nil {
		return
	}

	userID, err := a.repositoryUser.GetOrCreate(ctx, userData)
	if err != nil {
		return
	}

	jwtData := model.JWTAuthData{
		UserID: userID,
	}
	authToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtData)
	authJWTToken, err = authToken.SignedString(a.secretKey)
	return
}

func generateState() (string, error) {
	state := make([]byte, 64)
	_, err := rand.Read(state)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(state), nil
}

func (a *Auth) getHHData(ctx context.Context, code string) (*model.HHData, error) {
	params := url.Values{
		"client_id": {a.clientID},
		"client_secret": {a.clientSecret},
		"code": {code},
		"grant_type": {"authorization_code"},
		"redirect_uri": {a.redirectURI},
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://api.hh.ru/token", strings.NewReader(params.Encode()))
	if err != nil {
		return nil, domain.NewInternalServerError(domain.CodeInternalServerError, "unable to create request", err)
	}
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	request.Header.Set("HH-User-Agent", fmt.Sprintf("%s/%s (%s)", a.appName, a.appVersion, a.devContact))
	
	response, err := a.client.Do(request)
	if err != nil {
		return nil, domain.NewInternalServerError(domain.CodeInternalServerError, "unable to do request", err)
	}
	defer func ()  {
		closeErr := response.Body.Close()
		if closeErr != nil {
			err = domain.NewInternalServerError(domain.CodeInternalServerError, "unable to close response body", err, closeErr)
		}
	}()
	if response.StatusCode != http.StatusOK {
		return nil, domain.NewInternalServerError(domain.CodeInternalServerError, "response status is not ok")
	}

	var hhData model.HHData
	err = json.NewDecoder(response.Body).Decode(&hhData)
	if err != nil {
		return nil, domain.NewInternalServerError(domain.CodeInternalServerError, "unable to decode hh data", err)
	}
	return &hhData, nil
}

func (a *Auth) getUserData(ctx context.Context, accessToken string) (*model.UserData, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://api.hh.ru/me", nil)
	if err != nil {
		return nil, domain.NewInternalServerError(domain.CodeInternalServerError, "unable to create request", err)
	}
	request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", accessToken))
	request.Header.Set("HH-User-Agent", fmt.Sprintf("%s/%s (%s)", a.appName, a.appVersion, a.devContact))
	
	response, err := a.client.Do(request)
	if err != nil {
		return nil, domain.NewInternalServerError(domain.CodeInternalServerError, "unable to do request", err)
	}
	defer func ()  {
		closeErr := response.Body.Close()
		if closeErr != nil {
			err = domain.NewInternalServerError(domain.CodeInternalServerError, "unable to close response body", err, closeErr)
		}
	}()
	if response.StatusCode != http.StatusOK {
		return nil, domain.NewInternalServerError(domain.CodeInternalServerError, "response status is not ok")
	}

	var userData model.UserData
	err = json.NewDecoder(response.Body).Decode(&userData)
	if err != nil {
		return nil, domain.NewInternalServerError(domain.CodeInternalServerError, "unable to decode user data", err)
	}
	return &userData, nil
}
