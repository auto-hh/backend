package config

import "github.com/auto-hh/backend/pkg/config"

const (
	keyPort                  config.Key = "BACKEND_PORT"
	keyLogLevel              config.Key = "LOG_LEVEL"
	keySecretKey             config.Key = "SECRET_KEY"
	keyJWTExpirationDuration config.Key = "JWT_EXPIRATION_DURATION"
	keyStateExpirationDuration config.Key = "STATE_EXPIRATION_DURATION"
	keyClientID config.Key = "CLIENT_ID"
	keyClientSecret config.Key = "CLIENT_SECRET"
	keyRedirectURI config.Key = "REDIRECT_URI"
	keyAppName config.Key = "APP_NAME"
	keyAppVersion config.Key = "APP_VERSION"
	keyDevContact config.Key = "DEV_CONTACT"
	keySiteURL config.Key = "SITE_URL"
	keyLLMPath               config.Key = "LLM_PATH"
	KeyPostgresUser          config.Key = "POSTGRES_USER"
	//nolint:gosec
	KeyPostgresPassword config.Key = "POSTGRES_PASSWORD"
	KeyPostgresHost     config.Key = "POSTGRES_HOST"
	KeyPostgresPort     config.Key = "POSTGRES_PORT"
	KeyPostgresDatabase config.Key = "POSTGRES_DATABASE"
)
