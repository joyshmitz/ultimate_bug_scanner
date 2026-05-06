use axum::http::{header, HeaderMap, HeaderValue};
use tower_http::cors::CorsLayer;

struct Request;

impl Request {
    fn headers(&self) -> HeaderMap {
        HeaderMap::new()
    }
}

fn is_allowed_origin(origin: &str) -> bool {
    let allowed_origins = ["https://app.example.com", "https://admin.example.com"];
    allowed_origins.contains(&origin)
}

fn tower_http_allowlist_with_credentials() -> CorsLayer {
    let app_origin = HeaderValue::from_static("https://app.example.com");
    let admin_origin = HeaderValue::from_static("https://admin.example.com");
    CorsLayer::new()
        .allow_origin([app_origin, admin_origin])
        .allow_credentials(true)
}

fn reflected_origin_after_allowlist(request: Request, headers: &mut HeaderMap) {
    let origin = request
        .headers()
        .get("origin")
        .and_then(|value| value.to_str().ok())
        .unwrap_or_default();
    if !is_allowed_origin(origin) {
        return;
    }

    headers.insert(
        header::ACCESS_CONTROL_ALLOW_ORIGIN,
        match HeaderValue::from_str(origin) {
            Ok(value) => value,
            Err(_) => return,
        },
    );
    headers.insert(
        header::ACCESS_CONTROL_ALLOW_CREDENTIALS,
        HeaderValue::from_static("true"),
    );
    headers.insert(header::VARY, HeaderValue::from_static("Origin"));
}
