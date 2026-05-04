use std::collections::HashMap;
use std::env;

const JWT_SECRET: &str = "jwt_secret_from_rust_const_1234567890";
static STRIPE_SECRET_KEY: &str = "sk_live_rust_1234567890abcdef";

struct ProviderConfig {
    client_secret: String,
}

fn provider_config() -> ProviderConfig {
    ProviderConfig {
        client_secret: "oauth_client_secret_rust_1234567890".to_string(),
    }
}

fn session_secret() -> String {
    env::var("SESSION_SECRET").unwrap_or_else(|_| "session_secret_fallback_rust_123456".to_string())
}

fn webhook_secret() -> &'static str {
    option_env!("WEBHOOK_SECRET").unwrap_or("webhook_secret_fallback_rust_123456")
}

fn settings() -> HashMap<&'static str, String> {
    let mut settings = HashMap::new();
    settings.insert("API_KEY", "api_key_from_rust_dashboard_123456".to_string());
    let refresh_token = "refresh_token_from_rust_fixture_1234567890";
    settings.insert("REFRESH_TOKEN", refresh_token.to_string());
    settings
}
