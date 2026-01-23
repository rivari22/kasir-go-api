package main

import (
	"fmt"
	"kasir-api/handler"
	"log"
	"net/http"
)

func main() {
	fmt.Println("hello ini kasir API")
	port := ":8080"

	// get /categories
	http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		// handler here
		handler.GetCategories(w, r)
	})

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
