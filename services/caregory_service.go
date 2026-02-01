package services

import (
	"kasir-api/model"
	"kasir-api/repositories"
)

// CategoryService defines the interface for category business logic
type CategoryService interface {
	GetAllCategories() ([]model.Category, error)
	GetCategoryByID(id int) (*model.Category, error)
	CreateCategory(name, description string) (*model.Category, error)
	UpdateCategory(id int, name, description string) (*model.Category, error)
	DeleteCategory(id int) error
}

// categoryService implements CategoryService interface
type categoryService struct {
	repo repositories.CategoryRepository
}

// NewCategoryService creates a new category service instance
func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

// GetAllCategories retrieves all categories
func (s *categoryService) GetAllCategories() ([]model.Category, error) {
	return s.repo.GetAll()
}

// GetCategoryByID retrieves a category by its ID
func (s *categoryService) GetCategoryByID(id int) (*model.Category, error) {
	return s.repo.GetByID(id)
}

// CreateCategory creates a new category with validation
func (s *categoryService) CreateCategory(name, description string) (*model.Category, error) {
	// Business logic: additional validation or processing can be added here
	// For now, delegate to repository
	return s.repo.Create(name, description)
}

// UpdateCategory updates an existing category
func (s *categoryService) UpdateCategory(id int, name, description string) (*model.Category, error) {
	// Business logic: additional validation or processing can be added here
	// For now, delegate to repository
	return s.repo.Update(id, name, description)
}

// DeleteCategory deletes a category
func (s *categoryService) DeleteCategory(id int) error {
	return s.repo.Delete(id)
}
