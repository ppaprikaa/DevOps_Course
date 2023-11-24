package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, From Docker"))
	})

	srv := http.Server{
		Addr:    ":4000",
		Handler: mux,
	}

	log.Printf("Starting application at %s\n", srv.Addr)

	log.Fatal(srv.ListenAndServe())
}
