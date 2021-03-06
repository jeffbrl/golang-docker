package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// Set routing rules
	http.HandleFunc("/", handler)

	err := http.ListenAndServe(":8989", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "version 1")
}
