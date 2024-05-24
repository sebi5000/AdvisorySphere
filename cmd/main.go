package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	router.Get("/", indexHandler)
	router.Get("/employee_search", employeeSearchHandler)

	router.Post("/project_request", projectRequestHandler)
	http.ListenAndServe(":3000", router)
}
