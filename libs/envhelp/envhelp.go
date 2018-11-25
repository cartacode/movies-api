package envhelp

import (
	"os"
	"strconv"
)

// GetEnv ..
func GetEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// GetEnvInt ..
func GetEnvInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		v, _ := strconv.Atoi(value)
		return v
	}
	return fallback
}
