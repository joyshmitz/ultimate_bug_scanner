defmodule CleanCommandExecution do
  def git_status(repo_path) do
    System.cmd("git", ["-C", repo_path, "status", "--short"])
  end
end
