package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(".")))
	srv := &http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	srv.ListenAndServe()
}
