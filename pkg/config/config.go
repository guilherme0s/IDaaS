package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	Host, Port string
}

func Init() (Config, error) {
	cfg := Config{
		Host: GetEnvString("SERVER_HOST", "0.0.0.0"),
		Port: GetEnvString("SERVER_PORT", "8080"),
	}
	return cfg, nil
}

func GetEnvString(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return fallback
}

func GetEnvInt(key string, fallback int) int {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		return fallback
	}
	return i
}

func GetEnvDuration(key string, fallback time.Duration) time.Duration {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	d, err := time.ParseDuration(v)
	if err != nil {
		return fallback
	}
	return d
}
