package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s - %s %s\n", r.RemoteAddr, r.Proto, r.Method, r.URL)
	fmt.Fprintf(w, "Hello, World!\n")
}

func main() {
	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
