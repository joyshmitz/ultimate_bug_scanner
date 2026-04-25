# frozen_string_literal: true

require 'open3'

path = ARGV[0].to_s

# Safe argv array form prevents shell expansion
if !path.empty?
  system('rm', '--', path)
  pid = spawn('cat', '--', path)
  Process.wait(pid)
  stdout, _status = Open3.capture2('cat', '--', path)
  puts stdout
end
