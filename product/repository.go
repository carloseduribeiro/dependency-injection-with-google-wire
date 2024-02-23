package product

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// GetProductById returns the product with the given id.
// This Product Entity was not supposed to be returned. We should return a DTO instead.
// However, we will return it for now to keep example simple.
func (r *Repository) GetProductById(id int) (Product, error) {
	return Product{ID: id, Name: "Product Name"}, nil
}
