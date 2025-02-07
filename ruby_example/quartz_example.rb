# gem install quartz
# https://github.com/DavidHuie/quartz - utilizes the sharing of code between the 2 languages using net/rpc

require 'quartz'

client = Quartz::Client.new(file_path: 'adder.go')
puts client[:my_adder].call('Add', 'A' => 2, 'B' => 5)
