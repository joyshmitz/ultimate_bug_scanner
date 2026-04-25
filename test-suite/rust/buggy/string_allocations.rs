fn redundant_chain(input: &str) -> String {
    input.to_owned().to_string()
}

fn needless_format_literal() -> String {
    format!("static label")
}

fn needless_raw_format_literal() -> String {
    format!(r#"raw static label"#)
}
