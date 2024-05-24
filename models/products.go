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

	rows, err := db.Query("SELECT * FROM produtos")

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
