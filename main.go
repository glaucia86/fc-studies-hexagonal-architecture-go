/**
 * file: main.go
 * description: file responsible for the main application.
 * data: 22/01/2024
 * author: Glaucia Lemos <Twitter: @glaucia_lemos86>
 */

package main

import (
	"database/sql"

	db2 "github.com/glaucia86/fc-studies-hexagonal-architecture-go/adapters/db"
	"github.com/glaucia86/fc-studies-hexagonal-architecture-go/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db2.NewProductDb(db)

	productService := application.NewProductService(productDbAdapter)
	product, _ := productService.Create("Product 1", 10)

	productService.Enable(product)
}
