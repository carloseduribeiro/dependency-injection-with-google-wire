package product

import "database/sql"

type RepositoryInterface interface {
	GetProductById(id int) (Product, error)
}

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// GetProductById returns the product with the given id.
// This Product Entity was not supposed to be returned. We should return a DTO instead.
// However, we will return it for now to keep example simple.
func (r *ProductRepository) GetProductById(id int) (Product, error) {
	return Product{ID: id, Name: "Product Name"}, nil
}
