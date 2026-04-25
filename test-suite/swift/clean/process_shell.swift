import Foundation

func runFixedTool(_ path: String) throws {
    let process = Process()
    process.executableURL = URL(fileURLWithPath: "/usr/bin/stat")
    process.arguments = ["-f", "%z", path]
    try process.run()
}
