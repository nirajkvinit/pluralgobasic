package main

import (
	"net/http"

	"github.com/nirajkvinit/gowebservices/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
