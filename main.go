package main

import (
	"net/http"

	"github.com/JoaoPauloFontana/produtos/routes"
)

func main() {
	routes.Init()
	http.ListenAndServe(":8000", nil)
}
