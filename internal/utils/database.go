package utils

import (
	"auto-hh-backend/pkg/load_config"
	"time"
)

const (
	MAX_CONNS              load_config.ConfigKey = "MAX_CONNS"
	MAX_CONN_LIFE_TIME     load_config.ConfigKey = "MAX_CONN_LIFE_TIME"
	MAX_CONN_IDLE_TIME     load_config.ConfigKey = "MAX_CONN_IDLE_TIME"
	CONNECT_TIMEOUT        load_config.ConfigKey = "CONNECT_TIMEOUT"
	DefaultMaxConns        int32                 = 100
	DefaultMaxConnLifeTime                       = time.Hour
	DefaultMaxConnIdleTime                       = 30 * time.Minute
	DefaultConnectTimeout                        = 5 * time.Second
)
