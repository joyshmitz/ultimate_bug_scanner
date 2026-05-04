package security

import (
	"cmp"
	"os"
)

const jwtSecret = "jwt_secret_from_go_const_1234567890"

var stripeSecretKey = "sk_live_go_1234567890abcdef"

var authConfig = map[string]string{
	"SESSION_SECRET": "session_secret_from_go_map_1234567890",
	"WEBHOOK_SECRET": getenv("WEBHOOK_SECRET", "webhook_secret_fallback_go_123456"),
}

type ProviderConfig struct {
	ClientSecret string
}

func providerConfig() ProviderConfig {
	return ProviderConfig{
		ClientSecret: "oauth_client_secret_go_1234567890",
	}
}

func accessToken() string {
	return envOrDefault("ACCESS_TOKEN", "access_token_fallback_go_1234567890")
}

func jwtFallback() string {
	return cmp.Or(os.Getenv("JWT_SECRET"), "fallback_jwt_secret_go_123456")
}

func settings() map[string]string {
	settings := map[string]string{}
	settings["API_KEY"] = "api_key_from_go_dashboard_123456"
	return settings
}

func getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func envOrDefault(key, fallback string) string {
	return getenv(key, fallback)
}
