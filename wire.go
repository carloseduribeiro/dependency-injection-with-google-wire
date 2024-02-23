//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"
	"github.com/carloseduribeiro/dependency-injection-with-google-wire/product"

	"github.com/google/wire"
)

var setRepositoryDependency = wire.NewSet(
	product.NewProductRepository,
	wire.Bind(new(product.RepositoryInterface), new(*product.ProductRepository)),
)

func NewUseCase(db *sql.DB) *product.UseCase {
	wire.Build(
		setRepositoryDependency,
		product.NewProductUseCase,
	)
	return &product.UseCase{}
}
