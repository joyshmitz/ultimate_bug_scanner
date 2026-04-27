use std::process::Command;

fn insecure_command(user: &str) {
    Command::new("sh").arg("-c").arg(user).status().unwrap();
}

fn insecure_login_shell(user: &str) {
    Command::new("bash").args(["-lc", user]).status().unwrap();
}

fn insecure_windows_shell(user: &str) {
    Command::new("cmd").args(&["/C", user]).status().unwrap();
}

fn main() {
    let user = "rm -rf /";
    insecure_command(user);
    insecure_login_shell(user);
    insecure_windows_shell(user);
    println!("{:?}", std::env::var("API_KEY").unwrap_or("sk_live_123".into()));
}
