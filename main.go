package main

import (
	"kasir-api/config"
	"kasir-api/database"
	"kasir-api/handlers"
	"kasir-api/repositories"
	"kasir-api/services"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConnectionPort()
	db, err := database.InitDB(cfg.DBConn)
	if err != nil {
		log.Fatal("DB is not connected:", err)
	}
	defer db.Close()

	// Dependency Injection
	// Repository layer
	categoryRepo := repositories.NewCategoryRepository(db)
	// Service layer
	categoryService := services.NewCategoryService(categoryRepo)
	// Handler layer
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	// Setup routes
	mux := http.NewServeMux()

	// Category routes
	mux.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			categoryHandler.GetCategories(w, r)
		case http.MethodPost:
			categoryHandler.CreateCategory(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/categories/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			categoryHandler.GetCategoryByID(w, r)
		case http.MethodPut:
			categoryHandler.UpdateCategoryById(w, r)
		case http.MethodDelete:
			categoryHandler.DeleteCategoryById(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Apply CORS middleware
	handlerCors := config.CorsMiddleware(mux)

	// Get port from config or use default
	port := cfg.Port
	if port == "" {
		port = "8080"
	}
	port = ":" + port

	log.Println("Server running on port:", port)
	if err := http.ListenAndServe(port, handlerCors); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
