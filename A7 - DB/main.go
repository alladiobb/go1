package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, error := sql.Open("mysql", "root:root@tcp(localhost:3306)/golang")
	if error != nil {
		panic(error)
	}
	defer db.Close()

	product := NewProduct("Mouse", 49.99)
	product.Price = 100.99
	product.ID = "868e6d34-dcb9-454d-9e02-d97282b7ec5e"
	error = updateProduct(db, product)
	if error != nil {
		panic(error)
	}
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, error := db.Prepare("update  products set name = ? , price = ? where id = ?")
	if error != nil {
		panic(error)
	}
	defer stmt.Close()

	_, error = stmt.Exec(product.Name, product.Price, product.ID)
	if error != nil {
		panic(error)
	}
	return nil
}

func insertProduct(db *sql.DB, product *Product) error {
	stmt, error := db.Prepare("insert into products(id, name, price) values(?,?,?)")
	if error != nil {
		panic(error)
	}
	defer stmt.Close()

	_, error = stmt.Exec(product.ID, product.Name, product.Price)
	if error != nil {
		panic(error)
	}
	return nil
}
