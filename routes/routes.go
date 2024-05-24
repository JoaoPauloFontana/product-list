package routes

import (
	"net/http"

	"github.com/JoaoPauloFontana/produtos/controllers"
)

func Init() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
}
