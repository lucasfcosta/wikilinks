package api

import (
	"errors"
	"os"
	"strconv"
)

const defaultProtocol = "https"
const defaultAPIRoot = "en.wikipedia.org/w/api.php"
const defaultMaxParallelRequests = "10"

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// Config represents the configuration to access a particular Wikipedia API
type Config struct {
	Protocol, APIRoot   string
	MaxParallelRequests int
}

// NewConfigFromEnv creates an API config from environment variables using defaults if they don't exist
func NewConfigFromEnv() *Config {
	protocol := getEnv("WIKI_PROTOCOL", defaultProtocol)
	apiRoot := getEnv("WIKI_API_ROOT", defaultAPIRoot)
	maxParallelRequests, invalidParallelReqOption := strconv.Atoi(getEnv("WIKI_MAX_PARALLEL_REQUESTS", defaultMaxParallelRequests))

	if invalidParallelReqOption != nil {
		panic(errors.New("invalid WIKI_MAX_PARALLEL_REQUESTS option"))
	}

	return &Config{Protocol: protocol, APIRoot: apiRoot, MaxParallelRequests: maxParallelRequests}
}
