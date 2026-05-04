package main

import (
	"net/http"
)

type apiHandler struct{}

func (apiHandler) ServeHTTP(http.ResponseWriter, *http.Request) {}

func main() {
	mux := http.NewServeMux()
	srv := &http.Server{
		Handler: mux,
		Addr:    ":8080",
	}
	srv.ListenAndServe()
}
