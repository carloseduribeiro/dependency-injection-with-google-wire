//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"
	"github.com/carloseduribeiro/dependency-injection-with-google-wire/product"

	"github.com/google/wire"
)

func NewUseCase(db *sql.DB) *product.UseCase {
	wire.Build(
		product.NewProductRepository,
		product.NewProductUseCase,
	)
	return &product.UseCase{}
}
