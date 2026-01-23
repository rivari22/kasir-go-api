package handler

import (
	"encoding/json"
	"kasir-api/model"
	"net/http"
	"slices"
	"strconv"
	"sync"
)

// temporary store data, because no db
var (
	categories = []model.Category{}
	mu         sync.RWMutex
)

type genericReturn struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	defer mu.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(genericReturn{
		Data:    categories,
		Message: "success get categories",
	})
}

func GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	defer mu.RUnlock()

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

	// find category by id from slice categories
	var findCategory model.Category
	for i := 0; i < len(categories); i++ {
		if categories[i].ID == id {
			findCategory = categories[i]
			break
		}
	}

	// handle not found
	if findCategory.Name == "" && findCategory.ID == 0 {
		http.Error(w, "Category not found, Invalid category ID", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(genericReturn{
		Data:    findCategory,
		Message: "success get category by id",
	})
}

type categoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	// read and decode json body
	var categoryRequest categoryRequest
	if err := json.NewDecoder(r.Body).Decode(&categoryRequest); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// name validation, name is mandatory
	if categoryRequest.Name == "" {
		http.Error(w, "Name is empty", http.StatusBadRequest)
		return
	}

	// generate incremental ID from latest ID
	lastIndexCategories := len(categories) - 1
	latestId := 0
	if lastIndexCategories >= 0 {
		latestId = categories[lastIndexCategories].ID
	}
	newCategoryData := model.Category{
		ID:          latestId + 1,
		Name:        categoryRequest.Name,
		Description: categoryRequest.Description,
	}

	// append to categories var
	mu.Lock()
	categories = append(categories, newCategoryData)
	mu.Unlock()

	// return json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(genericReturn{
		Data:    newCategoryData.Name,
		Message: "success create new category",
	})
}

func UpdateCategoryById(w http.ResponseWriter, r *http.Request) {
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
	var categoryRequest categoryRequest
	if err := json.NewDecoder(r.Body).Decode(&categoryRequest); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// validate name is not empty
	if categoryRequest.Name == "" {
		http.Error(w, "Name should be not empty", http.StatusBadRequest)
		return
	}

	// make iteration and update data by same id
	isCategoryFound := false
	mu.Lock()
	for index := range categories {
		if categories[index].ID == id {
			categories[index] = model.Category{
				ID:          id,
				Name:        categoryRequest.Name,
				Description: categoryRequest.Description,
			}
			isCategoryFound = true
			break
		}
	}
	mu.Unlock()

	if !isCategoryFound {
		http.Error(w, "Category not found, invalid category ID", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(genericReturn{
		Data:    id,
		Message: "success update category by id",
	})
}

func DeleteCategoryById(w http.ResponseWriter, r *http.Request) {
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

	// find index category by id from slice categories
	indexToDelete := slices.IndexFunc(categories, func(category model.Category) bool {
		return category.ID == id
	})

	// handle not found
	if indexToDelete == -1 {
		http.Error(w, "Category not found, Invalid category ID", http.StatusBadRequest)
		return
	}

	// delete category by index
	mu.Lock()
	categories = slices.Delete(categories, indexToDelete, indexToDelete+1)
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(genericReturn{
		Data:    indexToDelete,
		Message: "success delete category by id",
	})
}
