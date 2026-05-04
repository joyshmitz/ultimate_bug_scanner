use std::collections::HashMap;
use std::env;

const PUBLIC_BASE_URL: &str = "https://example.com";

struct ProviderConfig {
    client_secret: String,
}

fn required_env(name: &str) -> Result<String, env::VarError> {
    env::var(name)
}

fn provider_config() -> Result<ProviderConfig, env::VarError> {
    Ok(ProviderConfig {
        client_secret: required_env("OAUTH_CLIENT_SECRET")?,
    })
}

fn session_secret() -> Result<String, env::VarError> {
    required_env("SESSION_SECRET")
}

fn webhook_secret() -> Result<String, env::VarError> {
    required_env("WEBHOOK_SECRET")
}

fn settings() -> Result<HashMap<&'static str, String>, env::VarError> {
    let mut settings = HashMap::new();
    settings.insert("API_KEY", required_env("API_KEY")?);
    let theme = "dark";
    settings.insert("THEME_VARIANT", theme.to_string());
    settings.insert("PUBLIC_BASE_URL", PUBLIC_BASE_URL.to_string());
    Ok(settings)
}
