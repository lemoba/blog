package controller

import (
	"html/template"
	"net/http"
)

func registerLoginRoutes() {
	http.HandleFunc("/admin/login", handlerLogin)
}

func handlerLogin(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("assets/templates/admin/admin_login.html")
	t.Execute(w, "hello")
}
