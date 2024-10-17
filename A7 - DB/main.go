package main

import (
	"database/sql"
	"fmt"

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
	db, error := sql.Open("mysql", "root:toor@tcp(localhost:3306)/golangg")
	if error != nil {
		panic(error)
	}
	defer db.Close()

	// product := NewProduct("Keyboard", 49.99)
	// // product.Price = 100.99
	// // product.ID = "868e6d34-dcb9-454d-9e02-d97282b7ec5e"
	// error = insertProduct(db, product)
	// if error != nil {
	// 	panic(error)
	// }

	// error = updateProduct(db, product)
	// if error != nil {
	// 	panic(error)
	// }

	deleteProduct(db, "07273862-ae57-4718-9584-4cd8df8ae86b")

	products, error := selectAllProducts(db)

	if error != nil {
		panic(error)
	}

	for _, product := range products {
		fmt.Printf("Product: %v, Preco: %.2f\n", product.Name, product.Price)
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

func selectOneProduct(db *sql.DB, id string) (*Product, error) {
	stmt, error := db.Prepare("select id, name, price from products where id = ?")

	if error != nil {
		return nil, error
	}

	defer stmt.Close()
	var product Product
	error = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)
	if error != nil {
		return nil, error
	}

	return &product, nil

}

func selectAllProducts(db *sql.DB) ([]Product, error) {
	// Não precisa fazer um STMT, olha só...
	// stmt, error := db.Prepare(select * from products)

	rows, error := db.Query("select * from products")

	if error != nil {
		return nil, error
	}

	defer db.Close()

	var products []Product
	for rows.Next() {
		var product Product
		error = rows.Scan(&product.ID, &product.Name, &product.Price)
		if error != nil {
			return nil, error
		}
		products = append(products, product)
	}
	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, error := db.Prepare("delete from products where id = ?")

	if error != nil {
		return error
	}

	defer stmt.Close()
	_, error = stmt.Exec(id)
	if error != nil {
		return error
	}
	return nil
}
