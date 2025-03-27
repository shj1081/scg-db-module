package config

import "time"

type DBConfig struct {
	DSN             string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

type ServerConfig struct {
	Port        string
	Environment string
}

type AuthConfig struct {
	ProxyURL string
}

type Config struct {
	DB     DBConfig
	Server ServerConfig
	Auth   AuthConfig
}
