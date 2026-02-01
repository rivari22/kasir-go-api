package handlers

import (
	"encoding/json"
	"kasir-api/model"
	"kasir-api/services"
	"net/http"
	"strconv"
)

// CategoryHandler handles HTTP requests for categories
type CategoryHandler struct {
	service services.CategoryService
}

// NewCategoryHandler creates a new category handler instance
func NewCategoryHandler(service services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

// GetCategories handles GET /categories - retrieves all categories
func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.service.GetAllCategories()
	if err != nil {
		http.Error(w, "Failed to retrieve categories: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.GenericReturn{
		Data:    categories,
		Message: "success get categories",
	})
}

// GetCategoryByID handles GET /categories/{id} - retrieves a category by ID
func (h *CategoryHandler) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	// get ID from path
	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, "Empty category ID", http.StatusBadRequest)
		return
	}

	// convert to int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	// get category from service
	category, err := h.service.GetCategoryByID(id)
	if err != nil {
		http.Error(w, "Failed to retrieve category: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if category == nil {
		http.Error(w, "Category not found, Invalid category ID", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.GenericReturn{
		Data:    category,
		Message: "success get category by id",
	})
}

// CreateCategory handles POST /categories - creates a new category
func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	// read and decode json body
	var req model.CategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// name validation, name is mandatory
	if req.Name == "" {
		http.Error(w, "Name is empty", http.StatusBadRequest)
		return
	}

	// create category via service
	category, err := h.service.CreateCategory(req.Name, req.Description)
	if err != nil {
		http.Error(w, "Failed to create category: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// return json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(model.GenericReturn{
		Data:    category,
		Message: "success create new category",
	})
}

// UpdateCategoryById handles PUT /categories/{id} - updates a category
func (h *CategoryHandler) UpdateCategoryById(w http.ResponseWriter, r *http.Request) {
	// get ID from path
	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, "Empty category ID", http.StatusBadRequest)
		return
	}

	// convert to int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	// get json from body req
	var req model.CategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// validate name is not empty
	if req.Name == "" {
		http.Error(w, "Name should be not empty", http.StatusBadRequest)
		return
	}

	// update category via service
	category, err := h.service.UpdateCategory(id, req.Name, req.Description)
	if err != nil {
		http.Error(w, "Failed to update category: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if category == nil {
		http.Error(w, "Category not found, invalid category ID", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.GenericReturn{
		Data:    category,
		Message: "success update category by id",
	})
}

// DeleteCategoryById handles DELETE /categories/{id} - deletes a category
func (h *CategoryHandler) DeleteCategoryById(w http.ResponseWriter, r *http.Request) {
	// get ID from path
	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, "Empty category ID", http.StatusBadRequest)
		return
	}

	// convert to int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	// delete category via service
	err = h.service.DeleteCategory(id)
	if err != nil {
		http.Error(w, "Failed to delete category: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.GenericReturn{
		Data:    id,
		Message: "success delete category by id",
	})
}
