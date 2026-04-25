fn interpolated_format(name: &str) -> String {
    format!("hello {name}")
}

fn escaped_brace_format() -> String {
    format!("{{status}}")
}

fn direct_string(input: &str) -> String {
    input.to_string()
}

fn documentation_mentions_only() -> &'static str {
    r#"
    These examples are documentation, not executable allocation smells:
    format!("static label")
    input.to_owned().to_string()
    "#
}

fn notes() {
    // format!("static label") should not count from comments.
    let _ = documentation_mentions_only();
}
