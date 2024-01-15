/**
 * file: adapters/db/product.go
 * description: file responsible for the db layer of the application.
 * data: 15/01/2024
 * author: Glaucia Lemos <Twitter: @glaucia_lemos86>
 */

package db

import (
	"database/sql"

	"github.com/glaucia86/fc-studies-hexagonal-architecture-go/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}

func (p *ProductDb) Get(id string) (application.ProductInterface, error) {
	var product application.Product
	stmt, err := p.db.Prepare("SELECT id, name, price, status FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}

	return &product, nil
}
