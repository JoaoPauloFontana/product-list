package routes

import (
	"net/http"

	"github.com/JoaoPauloFontana/produtos/controllers"
)

func Init() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/update", controllers.Update)
	http.HandleFunc("/delete", controllers.Delete)
}
