package handler

import (
	"encoding/json"
	"kasir-api/model"
	"net/http"
)

// temporary store data, because no db
var categories = []model.Category{}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}
