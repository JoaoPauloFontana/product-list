package routes

import (
	"net/http"

	"github.com/JoaoPauloFontana/produtos/controllers"
)

func Init() {
	http.HandleFunc("/", controllers.Index)
}
