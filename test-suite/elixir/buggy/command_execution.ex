defmodule BuggyCommandExecution do
  def list_with_shell(path) do
    System.cmd("sh", ["-c", "ls #{path}"])
  end

  def run_os_command(path) do
    :os.cmd(~c"ls #{path}")
  end

  def open_shell_port(path) do
    Port.open({:spawn, "sh -c 'cat #{path}'"}, [:binary])
  end
end
