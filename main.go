package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

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
	temp.ExecuteTemplate(w, "Index", nil)
}
