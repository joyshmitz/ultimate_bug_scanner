package security

import "os"

const publicBaseURL = "https://example.com"

var authConfig = map[string]string{
	"SESSION_SECRET": mustEnv("SESSION_SECRET"),
	"WEBHOOK_SECRET": mustEnv("WEBHOOK_SECRET"),
}

type ProviderConfig struct {
	ClientSecret string
}

func providerConfig() ProviderConfig {
	return ProviderConfig{
		ClientSecret: mustEnv("OAUTH_CLIENT_SECRET"),
	}
}

func accessToken() string {
	return mustEnv("ACCESS_TOKEN")
}

func jwtSecret() string {
	if value, ok := os.LookupEnv("JWT_SECRET"); ok && value != "" {
		return value
	}
	panic("JWT_SECRET must be configured")
}

func settings() map[string]string {
	settings := map[string]string{}
	settings["API_KEY"] = mustEnv("API_KEY")
	settings["THEME_VARIANT"] = "dark"
	settings["PUBLIC_BASE_URL"] = publicBaseURL
	return settings
}

func mustEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		panic("missing required environment variable: " + key)
	}
	return value
}
