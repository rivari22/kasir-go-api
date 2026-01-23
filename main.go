package main

import (
	"kasir-api/config"
	"kasir-api/handler"
	"log"
	"net/http"
)

func main() {
	port := ":8080"

	mux := http.NewServeMux()

	mux.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetCategories(w, r)
		case http.MethodPost:
			handler.CreateCategory(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/categories/{id}", func(w http.ResponseWriter, r *http.Request) {
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

	handlerCors := config.CorsMiddleware(mux)

	log.Println("Server running on port:", port)
	if err := http.ListenAndServe(port, handlerCors); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
