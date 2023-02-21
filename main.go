package main

import (
	controller "github.com/lemoba/blog/controller/admin"
	"net/http"
)

func main() {
	controller.RegisterAdminRoutes()
	http.ListenAndServe("localhost:8080", nil)
}
