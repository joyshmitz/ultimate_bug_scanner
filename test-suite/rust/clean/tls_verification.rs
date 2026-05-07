use native_tls::TlsConnector;
use openssl::ssl::{SslConnector, SslMethod, SslVerifyMode};
use reqwest::{Certificate, ClientBuilder};

fn reqwest_uses_default_tls() -> reqwest::Result<reqwest::Client> {
    ClientBuilder::new().https_only(true).build()
}

fn reqwest_uses_private_root(ca: Certificate) -> reqwest::Result<reqwest::Client> {
    ClientBuilder::new()
        .add_root_certificate(ca)
        .danger_accept_invalid_certs(false)
        .danger_accept_invalid_hostnames(false)
        .build()
}

fn native_tls_keeps_verification() -> Result<TlsConnector, native_tls::Error> {
    TlsConnector::builder().build()
}

fn openssl_keeps_peer_verification() -> Result<SslConnector, openssl::error::ErrorStack> {
    let mut builder = SslConnector::builder(SslMethod::tls())?;
    builder.set_verify(SslVerifyMode::PEER);
    Ok(builder.build())
}
