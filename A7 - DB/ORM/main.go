package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
)

type Product2 struct {
	ID    int `gorm:primaryKey`
	Name  string
	Price float64
}

func main() {
	dsn := "root:toor@tcp(localhost:3306)/golangg"
	db, error := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if error != nil {
		panic(error)
	}

	db.AutoMigrate(&Product2{})

	db.Create(&Product2{
		Name:  "Notebook",
		Price: 1000.00,
	})

	products := []Product2{
		{Name: "Notebook", Price: 1000},
		{Name: "Notebook2", Price: 2000},
		{Name: "Notebook3", Price: 3000},
	}

	db.Create(&products)
}
