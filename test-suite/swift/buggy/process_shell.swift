import Foundation

func runUserCommand(_ userInput: String) throws {
    let process = Process()
    process.executableURL = URL(fileURLWithPath: "/bin/sh")
    process.arguments = ["-c", "ls \(userInput)"]
    try process.run()
}

func runViaSystem(_ path: String) {
    system("ls \(path)")
}
