package main

import (
	"fmt"
	"net/http"
	"rest-api/middleware"

	"github.com/go-chi/chi/v5"

	"rest-api/api/routes"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.SetContentTypeMiddleware)

	routes.RegisterRoutes(r)

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}
}
