package main

import (
	"kasir-api/handler"
	"log"
	"net/http"
)

func main() {
	port := ":8080"

	http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetCategories(w, r)
		case http.MethodPost:
			handler.CreateCategory(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/categories/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetCategoryByID(w, r)
		case http.MethodPut:
			handler.UpdateCategoryById(w, r)
		case http.MethodDelete:
			handler.DeleteCategoryById(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
