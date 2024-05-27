package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	router.Get("/", indexHandler)
	//router.Get("/employee_search", employeeSearchHandler)

	router.Post("/project_request", projectRequestHandler)
	router.Get("/showExternalProfile", showExternalProfile)
	router.Post("/downloadExternalProfile", downloadExternalProfile)

	http.ListenAndServe(":8000", router)
}
