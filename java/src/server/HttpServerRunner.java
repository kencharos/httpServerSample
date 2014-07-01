package server;

import java.io.BufferedReader;
import java.io.File;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.OutputStream;
import java.net.InetSocketAddress;
import java.net.URLDecoder;
import java.nio.file.Files;
import java.util.HashMap;
import java.util.Map;
import java.util.regex.Pattern;

import com.sun.net.httpserver.HttpExchange;
import com.sun.net.httpserver.HttpHandler;
import com.sun.net.httpserver.HttpServer;

public class HttpServerRunner {

	/**
	 * @param args 0 - resource path, 1 - (Optional) port (default is 8080)
	 * @throws IOException
	 */
	public static void main(String[] args) throws IOException {
		
		final String path = args[0];
		
		int port = args.length == 1 ? 8080 : Integer.parseInt(args[1]);
		
		HttpServer server = HttpServer.create(new InetSocketAddress(port),0);
		
		server.createContext("/", new HttpHandler() {
			public void handle(HttpExchange ex) throws IOException {

				byte[] res = read(path, "welcom.html");
				
				ex.sendResponseHeaders(200, res.length);
				OutputStream os = ex.getResponseBody();
				os.write(res);
				os.close();
			}
		});
		
		server.createContext("/calc", new HttpHandler() {
			public void handle(HttpExchange ex) throws IOException {
				if (ex.getRequestMethod().equals("GET")) {
					byte[] res = read(path, "calc.html");
					
					ex.sendResponseHeaders(200, res.length);
					OutputStream os = ex.getResponseBody();
					os.write(res);
					os.close();
				} else {
					Map<String, String> param = param(ex.getRequestBody());
					
					String res;
					if (!isNumber(param.get("n1")) || ! isNumber(param.get("n2"))) {
						res = "Wrong input";
					} else {
						res = Integer.parseInt(param.get("n1")) + Integer.parseInt(param.get("n2")) +"";
					}
					
					String result = String.format(new String(read(path, "result.html")), res);
					
					byte[] b = result.getBytes();
					ex.sendResponseHeaders(200, b.length);
					OutputStream os = ex.getResponseBody();
					os.write(b);
					os.close();
				}
			}
		});
		
		server.setExecutor(null);
		server.start();
	}
	
	private static byte[] read(String path, String html) throws IOException {
		return Files.readAllBytes(new File(path, html).toPath());
	}
	
	private static final Pattern num = Pattern.compile("[+-]?[0-9]+");
	
	private static boolean isNumber(String n) {
		if (n == null) return false;
		return num.matcher(n).matches();
	}
	
	private static Map<String, String> param(InputStream is) throws IOException {
		Map<String, String> map = new HashMap<String, String>();
		try(BufferedReader br = new BufferedReader(new InputStreamReader(is))) {
			String line = null;
			while((line = br.readLine()) != null) {

				for(String item : URLDecoder.decode(line, "UTF-8").split("&")) {

					String[] s = item.split("=");
					map.put(s[0], s[1]);
				}
			}
		}
		return map;
	}
	
	
}
