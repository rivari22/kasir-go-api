package model

// GenericReturn defines the standard API response structure
type GenericReturn struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// CategoryRequest represents the request body for creating/updating a category
type CategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
