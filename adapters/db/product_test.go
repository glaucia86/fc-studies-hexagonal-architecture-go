/**
 * file: adapters/db/product_test.go
 * description: file responsible for the
 * data: 15/01/2024
 * author: Glaucia Lemos <Twitter: @glaucia_lemos86>
 */

package db

import (
	"database/sql"
	"log"
)

var Db *sql.DB

func SetUp() {
	// in this case, we are using the sqlite3 database and in memory
	Db, _ = sql.Open("sqlite3", ":memory:")
	CreateTable(Db)
	CreateProduct(Db)
}

func CreateTable(db *sql.DB) {
	table := `CREATE TABLE products ("id" string, "name" string, "price" float, "status" string);`

	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func CreateProduct(db *sql.DB) {
	insert := `INSERT INTO products VALUES ("abc", "Product Test",0,"disabled");`

	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}
