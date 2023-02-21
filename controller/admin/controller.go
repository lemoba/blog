package controller

import "net/http"

func RegisterAdminRoutes() {
	// 加载静态资源
	fs := http.FileServer(http.Dir("assets/bootstrap/"))
	http.Handle("/assets/bootstrap/", http.StripPrefix("/assets/bootstrap/", fs))

	registerLoginRoutes()
	registerIndexRoutes()
}
