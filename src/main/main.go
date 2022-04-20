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
	for name, values := range r.Header {
		for _, value := range values {
			fmt.Printf(name, value)
		}
	}
	fmt.Printf("\nHTTP request method: %s \n", r.Method)
	fmt.Printf("HTTP request body: %s \n", r.Body)
	fmt.Printf("-----------------------------------\n")

	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "<div><h1>Hello World</h1><a>fb.com</a></div>")
	default:
		fmt.Fprintf(w, "only get")
	}
}