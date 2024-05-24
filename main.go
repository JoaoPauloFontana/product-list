package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Qtd         int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := initDB()
	defer db.Close()
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func initDB() *sql.DB {
	connection := "user=postgres dbname=db_name password=sua_senha host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(fmt.Sprintf("Error opening database: %v", err))
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		panic(fmt.Sprintf("Error connecting to the database: %v", err))
	}

	return db
}

func index(w http.ResponseWriter, r *http.Request) {
	db := initDB()

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

	temp.ExecuteTemplate(w, "Index", products)

	defer db.Close()
}
