require 'webrick'


def num?(s) 
  return false if s == nil
  return /^[+-]?[0-9]+$/ =~ s
end


path = ARGV[0]
port = 3000
port = ARGV[1].to_i if ARGV.size > 1

server = WEBrick::HTTPServer.new({:Port => port})

server.mount_proc("/") {|req, res|
  filename = File.join(path, "welcom.html")
  open(filename) do |file|
    res.body = file.read
  end
  res.content_type = "text/html"
  #res.content_length = File.stat(filename).size
}

server.mount_proc("/calc") {|req, res|

  res.content_type = "text/html"
  if req.request_method == "GET"

    filename = File.join(path, "calc.html")
    open(filename) do |file|
      res.body = file.read
    end
    #res.content_length = File.stat(filename).size
  else 
    input = req.query
    result = "wrong input"
    if num?(input["n1"]) && num?(input["n2"])
      result = (input["n1"].to_i + input["n2"].to_i).to_s
    end
    filename = File.join(path, "result.html")
    content = nil
    open(filename) do |file|
      content = file.read
    end
    res.body = sprintf(content, result)
    #res.content_length = res.body.size
  end
}


trap(:INT){server.shutdown}
server.start