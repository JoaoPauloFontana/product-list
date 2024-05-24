package models

import (
	"fmt"

	"github.com/JoaoPauloFontana/produtos/db"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Qtd         int
}

func GetAllProduct() []Product {
	db := db.InitDB()

	rows, err := db.Query("SELECT * FROM produtos ORDER BY id ASC")

	if err != nil {
		panic(fmt.Sprintf("Error querying the database: %v", err))
	}

	p := Product{}
	products := []Product{}

	for rows.Next() {
		err = rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Qtd)

		if err != nil {
			panic(fmt.Sprintf("Error scanning the database: %v", err))
		}

		products = append(products, p)
	}

	defer db.Close()

	return products
}

func InsertProduct(name string, description string, price float64, qtd int) {
	db := db.InitDB()

	insertData, err := db.Prepare("INSERT INTO produtos (name, description, price, qtd) VALUES ($1, $2, $3, $4)")

	if err != nil {
		panic(fmt.Sprintf("Error preparing the statement: %v", err))
	}

	_, err = insertData.Exec(name, description, price, qtd)

	if err != nil {
		panic(fmt.Sprintf("Error executing the statement: %v", err))
	}

	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.InitDB()

	deleteData, err := db.Prepare("DELETE FROM produtos WHERE id = $1")

	if err != nil {
		panic(fmt.Sprintf("Error preparing the statement: %v", err))
	}

	_, err = deleteData.Exec(id)

	if err != nil {
		panic(fmt.Sprintf("Error executing the statement: %v", err))
	}

	defer db.Close()
}

func GetProduct(id string) Product {
	db := db.InitDB()

	row := db.QueryRow("SELECT * FROM produtos WHERE id = $1", id)

	p := Product{}

	err := row.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Qtd)

	if err != nil {
		panic(fmt.Sprintf("Error scanning the database: %v", err))
	}

	defer db.Close()

	return p
}

func UpdateProduct(id int, name string, description string, price float64, qtd int) {
	db := db.InitDB()

	updateData, err := db.Prepare("UPDATE produtos SET name = $1, description = $2, price = $3, qtd = $4 WHERE id = $5")

	if err != nil {
		panic(fmt.Sprintf("Error preparing the statement: %v", err))
	}

	_, err = updateData.Exec(name, description, price, qtd, id)

	if err != nil {
		panic(fmt.Sprintf("Error executing the statement: %v", err))
	}

	defer db.Close()
}
