package webserver

func Web_server_main() {
	server := NewServer(":3000")
	server.Handle("/", server.AddMiddleware(HandleRoot, Logging()))
	server.Handle("/api", server.AddMiddleware(HandleApi, Logging(), CheckAuth()))
	server.Listen()
}
