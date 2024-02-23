package main

import (
	"database/sql"
	"fmt"

	"github.com/carloseduribeiro/dependency-injection-with-google-wire/product"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := product.NewProductRepository(db)
	useCase := product.NewProductUseCase(repository)

	p, err := useCase.Execute(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(p.Name)
}
