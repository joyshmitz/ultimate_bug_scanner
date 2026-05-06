use axum::http::{header, HeaderMap, HeaderValue};
use tower_http::cors::{Any, CorsLayer};

struct Request;

impl Request {
    fn headers(&self) -> HeaderMap {
        HeaderMap::new()
    }
}

fn tower_http_any_origin_with_credentials() -> CorsLayer {
    CorsLayer::new()
        .allow_origin(Any)
        .allow_credentials(true)
}

fn actix_any_origin_with_credentials() {
    let _cors = actix_cors::Cors::default()
        .allow_any_origin()
        .supports_credentials();
}

fn wildcard_header_with_credentials(headers: &mut HeaderMap) {
    headers.insert(
        header::ACCESS_CONTROL_ALLOW_ORIGIN,
        HeaderValue::from_static("*"),
    );
    headers.insert(
        header::ACCESS_CONTROL_ALLOW_CREDENTIALS,
        HeaderValue::from_static("true"),
    );
}

fn reflected_origin_with_credentials(request: Request, headers: &mut HeaderMap) {
    let origin = request.headers().get("origin").unwrap().clone();
    headers.insert(header::ACCESS_CONTROL_ALLOW_ORIGIN, origin);
    headers.insert(
        header::ACCESS_CONTROL_ALLOW_CREDENTIALS,
        HeaderValue::from_static("true"),
    );
}
