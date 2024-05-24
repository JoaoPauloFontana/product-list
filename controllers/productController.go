package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/JoaoPauloFontana/produtos/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.GetAllProduct()

	temp.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		qtd := r.FormValue("qtd")

		convertedPrice, err := strconv.ParseFloat(price, 64)

		if err != nil {
			panic("Error converting price: " + err.Error())
		}

		convertedQtd, err := strconv.Atoi(qtd)

		if err != nil {
			panic("Error converting qtd: " + err.Error())
		}

		models.InsertProduct(name, description, convertedPrice, convertedQtd)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	models.DeleteProduct(id)

	http.Redirect(w, r, "/", 301)
}
