package controller

import (
	"html/template"
	"net/http"
)

func registerIndexRoutes() {
	http.HandleFunc("/admin/index", handlerIndex)
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("assets/templates/admin/admin_index.html")
	t.Execute(w, "hello")
}
