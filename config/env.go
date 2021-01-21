package config

import "os"

func Env(key string, defaultValue string) string {
	if len(os.Getenv(key)) > 0 {
		return os.Getenv(key)
	}

	return defaultValue
}
