package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("POST /api/create", nil)
	mux.Handle("POST /api/update", nil)
	mux.Handle("POST /api/select", nil)
	mux.Handle("POST /api/delete", nil)
	mux.Handle("POST /api/list", nil)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
		//add after
	}

	log.Fatal(server.ListenAndServe())
}
