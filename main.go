package main

import (
	"html/template"
	"net/http"

	"github.com/JoaoPauloFontana/produtos/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := models.GetAllProduct()
	temp.ExecuteTemplate(w, "Index", products)
}
