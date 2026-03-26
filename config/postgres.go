package config

import (
	"fmt"
	"net"
)

type PostgresConfig struct {
	user     string
	password string
	host     string
	port     string
	database string
}

func LoadPostgresConfig() *PostgresConfig {
	user := KeyPostgresUser.GetValueDefault("user")
	password := KeyPostgresPassword.GetValueDefault("password")
	host := KeyPostgresHost.GetValueDefault("database")

	port := KeyPostgresPort.GetValueDefault("5432")

	database := KeyPostgresDatabase.GetValueDefault("database")

	return &PostgresConfig{
		user:     user,
		password: password,
		host:     host,
		port:     port,
		database: database,
	}
}

func (c *PostgresConfig) DSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		c.user,
		c.password,
		net.JoinHostPort(c.host, c.port),
		c.database,
	)
}
