package main

import (
	"fmt"
	"net/http"
	"log"
)




func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")
}


func echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Echo: Hello World!\n")
}



func main(){
	mux := http.NewServeMux()

    mux.HandleFunc("/hello", hello	)
    mux.HandleFunc("/echo", echo)


	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,

	}

	log.Println("Starting server on :8080")
	log.Fatal(server.ListenAndServe())
}