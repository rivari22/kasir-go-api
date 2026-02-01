package main

import (
	"fmt"
	"kasir-api/config"
	"kasir-api/database"
	"log"
)

// old handler
type Printer struct {
	prefix string
	subfix string
}

func NewPrinter(prefix string, subfix string) *Printer {
	return &Printer{prefix: prefix, subfix: subfix}
}

func (p *Printer) Print(name string) {
	fmt.Println(p.prefix, name, p.subfix)
}

func main() {
	config := config.LoadConnectionPort()
	fmt.Println(config.DBConn)
	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal(err, " DB is not connected")
	}

	defer db.Close()

	// TODO: MOVE CODE BELOW
	// initPrinter := NewPrinter("Halo", "Selamat Belajar")
	// initPrinter.Print("Rivari")
	// fmt.Println(initPrinter)
	// port := ":8080"

	// mux := http.NewServeMux()

	// mux.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
	// 	switch r.Method {
	// 	case http.MethodGet:
	// 		handler.GetCategories(w, r)
	// 	case http.MethodPost:
	// 		handler.CreateCategory(w, r)
	// 	default:
	// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	// 	}
	// })

	// mux.HandleFunc("/categories/{id}", func(w http.ResponseWriter, r *http.Request) {
	// 	switch r.Method {
	// 	case http.MethodGet:
	// 		handler.GetCategoryByID(w, r)
	// 	case http.MethodPut:
	// 		handler.UpdateCategoryById(w, r)
	// 	case http.MethodDelete:
	// 		handler.DeleteCategoryById(w, r)
	// 	default:
	// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	// 	}
	// })

	// handlerCors := config.CorsMiddleware(mux)

	// log.Println("Server running on port:", port)
	// if err := http.ListenAndServe(port, handlerCors); err != nil {
	// 	log.Fatal("Error starting server:", err)
	// }
}
