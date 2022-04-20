package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleHttp)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handleHttp(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("HTTP request headers\n")
	
}