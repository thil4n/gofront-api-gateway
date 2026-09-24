package main

import (
	"fmt"
	"net/http"
	"log"
)



func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")
}


func main(){
	http.HandleFunc("/hello", hello)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}