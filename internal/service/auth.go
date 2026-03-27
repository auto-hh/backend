package service

import (
	"crypto/rand"
	"encoding/base64"
	"net/url"

	"github.com/auto-hh/backend/internal/domain"
	"github.com/auto-hh/backend/internal/model"
	"github.com/auto-hh/backend/internal/repository"
	"github.com/golang-jwt/jwt/v5"
)

type Auth struct {
	repositoryUser repository.IUser
	secretKey []byte
	clientID string
	redirectURI string
}

func NewAuth(repositoryUser repository.IUser, secretKey []byte, clientID string, redirectURI string) *Auth {
	return &Auth{
		repositoryUser: repositoryUser,
		secretKey: secretKey,
		clientID: clientID,
		redirectURI: redirectURI,
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

func generateState() (string, error) {
	state := make([]byte, 64)
	_, err := rand.Read(state)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(state), nil
}
