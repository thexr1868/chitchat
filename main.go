package main

import "net/http"

func main() {
	p("ChitChat", version(), "started at", config.Address)

	// handle static assets
	mux := http.NewServeMux() // multiplexer that redirects a request to a handler
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)

	// starting up the server
	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    0,
		WriteTimeout:   0,
		MaxHeaderBytes: 0,
	}
	server.ListenAndServe()
}
