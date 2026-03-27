package config

import (
	"fmt"
	"log/slog"
	"strconv"
	"time"
)

type Config struct {
	LogLevel              slog.Level
	port                  int
	SecretKey             []byte
	JWTExpirationDuration time.Duration
	StateExpirationDuration time.Duration
	ClientID string
	ClientSecret string
	RedirectURI string
	AppName string
	AppVersion string
	DevContact string
	SiteURL string
	LLMPath               string
	postgresConfig        *PostgresConfig
}

func LoadConfig() (*Config, error) {
	var logLevel slog.Level

	err := logLevel.UnmarshalText([]byte(keyLogLevel.GetValueDefault("info")))
	if err != nil {
		return nil, err
	}

	port, err := strconv.Atoi(keyPort.GetValueDefault("8080"))
	if err != nil {
		return nil, err
	}

	secretKey := []byte(keySecretKey.GetValueDefault("secret-key"))

	jwtExpirationDuration, err := strconv.Atoi(keyJWTExpirationDuration.GetValueDefault("24"))
	if err != nil {
		return nil, err
	}
	stateExpirationDuration, err := strconv.Atoi(keyStateExpirationDuration.GetValueDefault("10"))
	if err != nil {
		return nil, err
	}

	clientID := keyClientID.GetValueDefault("client_id")
	clientSecret := keyClientSecret.GetValueDefault("client_secret")
	redirectURI := keyRedirectURI.GetValueDefault("redirect_uri")
	appName := keyRedirectURI.GetValueDefault("app_name")
	appVersion := keyRedirectURI.GetValueDefault("app_version")
	devContact := keyRedirectURI.GetValueDefault("dev_contact")
	siteURL := keySiteURL.GetValueDefault("https://localhost/")

	llmPath := keyLLMPath.GetValue()

	postgresConfig := LoadPostgresConfig()

	config := &Config{
		LogLevel:              logLevel,
		port:                  port,
		SecretKey:             secretKey,
		JWTExpirationDuration: time.Duration(jwtExpirationDuration) * time.Hour,
		StateExpirationDuration: time.Duration(stateExpirationDuration) * time.Minute,
		ClientID: clientID,
		ClientSecret: clientSecret,
		RedirectURI: redirectURI,
		AppName: appName,
		AppVersion: appVersion,
		DevContact: devContact,
		SiteURL: siteURL,
		LLMPath:               llmPath,
		postgresConfig:        postgresConfig,
	}

	return config, nil
}

func (c *Config) PostgresDSN() string {
	return c.postgresConfig.DSN()
}

func (c *Config) Address() string {
	return fmt.Sprintf(":%d", c.port)
}
