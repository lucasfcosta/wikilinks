package api

import "os"

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// Config represents the configuration to access a particular Wikipedia API
type Config struct {
	Protocol, APIRoot string
}

// NewConfigFromEnv creates an API config from environment variables using defaults if they don't exist
func NewConfigFromEnv() *Config {
	var protocol = getEnv("WIKI_PROTOCOL", "https")
	var apiRoot = getEnv("WIKI_API_ROOT", "en.wikipedia.org/w/api.php")
	return &Config{Protocol: protocol, APIRoot: apiRoot}
}
