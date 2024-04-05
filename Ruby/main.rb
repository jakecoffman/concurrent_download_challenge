# frozen_string_literal: true

require 'net/http'

urls_to_load = %w[http://google.com http://python.org http://ruby-lang.org http://golang.org]

start_time = Time.now
limit = 0.5
queue = Queue.new

def fetch(url)
  loop do
    uri = URI(url)
    response = Net::HTTP.get_response(uri)
    if response.code == '301' || response.code == '302'
      url = response['location']
      next
    end
    return response
  end
end

threads = urls_to_load.map do |url|
  Thread.new do
    Timeout.timeout(limit) do
      fetch_start_time = Time.now
      response = fetch(url)
      queue << "#{url} #{response.body.to_s.size} #{Time.now - fetch_start_time}"
    end
  rescue Timeout::Error
    # Ignored
  end
end

threads.each(&:join)

queue.size.times do
  puts queue.pop
end

puts "Time elapsed: #{Time.now - start_time}"
