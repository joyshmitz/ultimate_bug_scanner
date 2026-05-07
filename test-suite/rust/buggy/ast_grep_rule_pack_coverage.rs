// Intentionally dense UBS rule-pack fixture.
// This file is parsed by ast-grep but is not compiled; unresolved external
// crates/types are deliberate so each generated Rust rule has a compact sample.

use std::cell::RefCell;
use std::convert::TryInto;
use std::ffi::CStr;
use std::mem::{self, MaybeUninit};
use std::rc::Rc;
use std::str;
use std::sync::{Arc, Mutex, RwLock};
use std::time::{Duration, Instant};

struct RiskyDrop {
    value: Option<String>,
}

impl Drop for RiskyDrop {
    fn drop(&mut self) {
        let _ = self.value.take().unwrap();
        panic!("panic in drop");
    }
}

struct NotThreadSafe(*mut u8);

unsafe impl Send for NotThreadSafe {}
unsafe impl Sync for NotThreadSafe {}

async fn async_hazards(
    std_mutex: Mutex<i32>,
    std_rwlock: RwLock<i32>,
    tokio_mutex: tokio::sync::Mutex<i32>,
    tokio_rwlock: tokio::sync::RwLock<i32>,
    futures: Vec<impl std::future::Future<Output = ()>>,
) {
    let _std_guard = std_mutex.lock().unwrap();
    async_work().await;

    let _std_expect_guard = std_mutex.lock().expect("lock");
    async_work().await;

    let _std_read_guard = std_rwlock.read().unwrap();
    async_work().await;

    let _tokio_guard = tokio_mutex.lock().await;
    async_work().await;

    let _tokio_read_guard = tokio_rwlock.read().await;
    async_work().await;

    for fut in futures {
        fut.await;
    }

    futures::executor::block_on(async_work());
    tokio::runtime::Runtime::block_on(async_work());
    std::fs::read("config.txt").unwrap();
    std::thread::sleep(Duration::from_millis(1));
    std::thread::spawn(|| println!("thread in async"));
    tokio::spawn(async { async_work().await });
    tokio::spawn(async move { async_work().await });
    let _task = tokio::spawn(async move { async_work().await });
    tokio::task::block_in_place(|| std::fs::read_to_string("blocking.txt").unwrap());
}

async fn async_work() {}

fn macro_and_panic_hazards(left: i32, right: i32) {
    assert!(left > 0);
    assert_eq!(left, right);
    assert_ne!(left, right + 1);
    debug_assert!(left < right);
    debug_assert_eq!(left, right);
    debug_assert_ne!(left, right + 1);
    dbg!(left);
    eprintln!("debug output");
    println!("debug output");
    panic!("explicit panic");
    todo!("placeholder");
    unimplemented!("missing branch");
    unreachable!("should be impossible");
}

fn ownership_and_iteration(items: &[String], input: &str) -> Vec<String> {
    let mut out = Vec::new();

    for item in items {
        out.push(item.clone());
    }

    for item in items.iter().cloned().collect::<Vec<String>>() {
        out.push(item);
    }

    let mapped: Vec<String> = items.iter().map(|item| item.clone()).collect();
    out.extend(mapped);
    let _ = items.iter().nth(0);
    let _ = input.chars().nth(3);
    let _ = input.chars().nth(3).unwrap();
    let _ = input.to_owned().to_string();
    out
}

fn allocation_and_regex_hazards(inputs: &[&str]) {
    for input in inputs {
        let _ = input.clone();
        let re = regex::Regex::new(r"^[a-z]+$").unwrap();
        let _ = re.is_match(input);
    }

    while inputs.len() > 1 {
        let _ = regex::Regex::new(r"\d+");
        break;
    }

    let _ = regex::Regex::new(r"^\w+$").unwrap();
    let _ = format!("static label");
}

fn unsafe_memory_hazards(bytes: &[u8], index: usize, ptr: *const u8) {
    unsafe {
        let _ = bytes.get_unchecked(index);
        let _ = bytes.get_unchecked_mut(index);
        let _ = std::slice::from_raw_parts(ptr, bytes.len());
        let _ = std::slice::from_raw_parts_mut(ptr as *mut u8, bytes.len());
        let _: usize = std::mem::transmute(ptr);
        let _ = std::mem::uninitialized::<String>();
        let _ = std::mem::zeroed::<String>();
        let maybe: MaybeUninit<String> = MaybeUninit::uninit();
        let _ = maybe.assume_init();
        let _ = std::hint::unreachable_unchecked();
        let _ = str::from_utf8_unchecked(bytes);
        let _ = String::from_utf8_unchecked(bytes.to_vec());
        let _ = CStr::from_bytes_with_nul_unchecked(bytes);
        let _ = Some(1).unwrap_unchecked();
    }

    let _ = bytes as *const [u8];
    let _ = ptr as *mut u8;
    mem::forget(bytes.to_vec());
}

fn conversion_and_numeric_hazards(values: &[u64], input: &str) {
    let _ = values.len() as u8;
    let _ = values.iter().count() as i16;
    let _ = values.len() as f64;
    let _ = input.parse::<u64>().unwrap();
    let _ = input.parse::<u64>().expect("number");
    let _ = input.parse::<f64>().unwrap_or(0.0);
    let _: [u8; 4] = values.try_into().unwrap();
    let _ = String::from_utf8(values.iter().map(|v| *v as u8).collect()).unwrap();
    let _ = str::from_utf8(b"\xff").expect("utf8");
    let _ = 1_i64.wrapping_neg();
    let _ = -(values.len() as i64);
    let _ = 1_u64.wrapping_add(values.len() as u64);
    let _ = 1_u64.wrapping_sub(values.len() as u64);
    let _ = 1_u64.wrapping_mul(values.len() as u64);
    let _ = Instant::now().elapsed();
    let _ = Instant::now() - Duration::from_secs(1);
}

fn security_api_hazards(data: &[u8], user_command: &str) {
    let _ = std::process::Command::new("sh").arg("-c").arg(user_command);
    let _ = std::process::Command::new("sh").args(["-c", user_command]);
    let _ = std::env::var("TOKEN").unwrap();
    let _ = serde_json::from_str::<serde_json::Value>("{}").unwrap();
    let _ = md5::compute(data);
    let _ = sha1::Sha1::new();
    let _ = openssl::hash::MessageDigest::md5();
    let _ = openssl::ssl::SslVerifyMode::NONE;
    let _ = native_tls::TlsConnector::builder().danger_accept_invalid_certs(true);
    let _ = native_tls::TlsConnector::builder().danger_accept_invalid_hostnames(true);
    let _ = reqwest::ClientBuilder::new().danger_accept_invalid_certs(true);
    let _ = reqwest::ClientBuilder::new().danger_accept_invalid_hostnames(true);
    let _ = Nonce::from_slice(data);
    let _ = GenericArray::from_slice(data);
    let _ = Key::from_slice(data);
    let _ = "http://example.invalid/api";
}

fn resource_hazards() {
    std::thread::spawn(|| println!("discarded join handle"));
    let _handle = std::thread::spawn(|| println!("unjoined thread"));
    std::fs::write("state.txt", b"not atomic").unwrap();
}

fn container_type_hazards() {
    let _shared: Option<Arc<Mutex<Vec<String>>>> = None;
    let _local: Option<Rc<RefCell<Vec<String>>>> = None;
}
