package controllers

import (
	"html/template"
	"net/http"

	"github.com/JoaoPauloFontana/produtos/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.GetAllProduct()

	temp.ExecuteTemplate(w, "Index", products)
}
