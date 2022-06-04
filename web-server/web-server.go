package webserver

func Web_server_main() {
	server := NewServer(":3000")
	server.Handle("/", HandleRoot)
	server.Handle("/api", server.AddMiddleware(HandleApi, CheckAuth()))
	server.Listen()
}
