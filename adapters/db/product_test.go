/**
 * file: adapters/db/product_test.go
 * description: file responsible for the
 * data: 16/01/2024
 * author: Glaucia Lemos <Twitter: @glaucia_lemos86>
 */

package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/glaucia86/fc-studies-hexagonal-architecture-go/adapters/db"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func SetUp() {
	// in this case, we are using the sqlite3 database and in memory
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products ("id" string, "name" string, "price" float, "status" string);`

	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products VALUES ("abc", "Product Test",0,"disabled");`

	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	SetUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)

	product, err := productDb.Get("abc")
	require.Nil(t, err, "Error when trying to get a product")
	require.Equal(t, "Product Test", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}
