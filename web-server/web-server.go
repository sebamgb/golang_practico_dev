package webserver

func Web_server_main() {
	server := NewServer(":3000")
	server.Handle("GET", "/", server.AddMiddleware(HandleRoot, Logging()))
	server.Handle("POST", "/create", PostRequest)
	server.Handle("POST", "/api", server.AddMiddleware(HandleApi, Logging(), CheckAuth()))
	server.Listen()
}
