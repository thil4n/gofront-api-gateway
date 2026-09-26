package main

import (
	"fmt"
	"net/http"
	"log"
)


func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK!\n")
}

func echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Method: %s\n", r.Method)
	fmt.Fprintf(w, "Path: %s\n", r.URL.Path)
	fmt.Fprintf(w, "Query: %s\n", r.URL.RawQuery)
	fmt.Fprintf(w, "Remote Address: %s\n", r.RemoteAddr)
}

func main(){
	mux := http.NewServeMux()

    mux.HandleFunc("/health", health	)
    mux.HandleFunc("/echo", echo)


	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,

	}

	log.Println("Starting server on :8080")
	log.Fatal(server.ListenAndServe())
}