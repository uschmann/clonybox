package config

import "os"

type Config struct {
	DatabaseFile string
}

func FromEnv() *Config {
	return &Config{
		DatabaseFile: getEnv("DB_FILE", "database.db"),
	}
}

func getEnv(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)

	if ok {
		return value
	}

	return defaultValue
}
