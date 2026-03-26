package config

import (
	"log/slog"
	"strconv"
	"time"
)

type Config struct {
	logLevel          slog.Level
	port              int
	secretKey         []byte
	jwtExpirationDuration time.Duration
	postgresConfig    *PostgresConfig
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

	postgresConfig := LoadPostgresConfig()

	config := &Config{
		logLevel:          logLevel,
		port:              port,
		secretKey:         secretKey,
		jwtExpirationDuration: time.Duration(jwtExpirationDuration) * time.Hour,
		postgresConfig:    postgresConfig,
	}

	return config, nil
}

func (c *Config) LogLevel() slog.Level {
	return c.logLevel
}

func (c *Config) Port() int {
	return c.port
}

func (c *Config) GetSecretKey() []byte {
	return c.secretKey
}

func (c *Config) JWTExpirationDuration() time.Duration {
	return c.jwtExpirationDuration
}

func (c *Config) PostgresDSN() string {
	return c.postgresConfig.DSN()
}
