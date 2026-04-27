use std::process::Command;

fn safe_command(user: &str) {
    Command::new("ls").arg(user).status().unwrap();
}

fn safe_args_command(user: &str) {
    let _ = Command::new("grep").args(["--", user, "Cargo.toml"]).status();
}

fn main() {
    let user = "docs";
    safe_command(user);
    safe_args_command(user);
}
