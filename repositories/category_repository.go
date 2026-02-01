package repositories

import (
	"database/sql"
	"kasir-api/model"
)

// CategoryRepository defines the interface for category database operations
type CategoryRepository interface {
	GetAll() ([]model.Category, error)
	GetByID(id int) (*model.Category, error)
	Create(name, description string) (*model.Category, error)
	Update(id int, name, description string) (*model.Category, error)
	Delete(id int) error
}

// categoryRepository implements CategoryRepository interface
type categoryRepository struct {
	db *sql.DB
}

// NewCategoryRepository creates a new category repository instance
func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

// GetAll retrieves all categories from the database
func (r *categoryRepository) GetAll() ([]model.Category, error) {
	rows, err := r.db.Query("SELECT id, name, description FROM category ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []model.Category
	for rows.Next() {
		var category model.Category
		if err := rows.Scan(&category.ID, &category.Name, &category.Description); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

// GetByID retrieves a category by its ID
func (r *categoryRepository) GetByID(id int) (*model.Category, error) {
	var category model.Category
	err := r.db.QueryRow("SELECT id, name, description FROM category WHERE id = $1", id).
		Scan(&category.ID, &category.Name, &category.Description)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Category not found
		}
		return nil, err
	}

	return &category, nil
}

// Create inserts a new category into the database
func (r *categoryRepository) Create(name, description string) (*model.Category, error) {
	var category model.Category
	err := r.db.QueryRow(
		"INSERT INTO category (name, description) VALUES ($1, $2) RETURNING id, name, description",
		name, description,
	).Scan(&category.ID, &category.Name, &category.Description)

	if err != nil {
		return nil, err
	}

	return &category, nil
}

// Update updates an existing category in the database
func (r *categoryRepository) Update(id int, name, description string) (*model.Category, error) {
	var category model.Category
	err := r.db.QueryRow(
		"UPDATE category SET name = $1, description = $2 WHERE id = $3 RETURNING id, name, description",
		name, description, id,
	).Scan(&category.ID, &category.Name, &category.Description)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Category not found
		}
		return nil, err
	}

	return &category, nil
}

// Delete removes a category from the database
func (r *categoryRepository) Delete(id int) error {
	result, err := r.db.Exec("DELETE FROM category WHERE id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows // Category not found
	}

	return nil
}
