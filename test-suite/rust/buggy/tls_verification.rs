use native_tls::TlsConnector;
use openssl::ssl::{SslConnector, SslMethod, SslVerifyMode};
use reqwest::ClientBuilder;

fn reqwest_accepts_invalid_certificates() -> reqwest::Result<reqwest::Client> {
    ClientBuilder::new()
        .danger_accept_invalid_certs(true)
        .build()
}

fn reqwest_accepts_invalid_hostnames() -> reqwest::Result<reqwest::Client> {
    ClientBuilder::new()
        .danger_accept_invalid_hostnames(true)
        .build()
}

fn native_tls_accepts_invalid_hostnames() -> Result<TlsConnector, native_tls::Error> {
    TlsConnector::builder()
        .danger_accept_invalid_hostnames(true)
        .build()
}

fn openssl_accepts_any_certificate() -> Result<SslConnector, openssl::error::ErrorStack> {
    let mut builder = SslConnector::builder(SslMethod::tls())?;
    builder.set_verify(SslVerifyMode::NONE);
    Ok(builder.build())
}
