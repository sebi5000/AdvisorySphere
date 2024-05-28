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
	router.Get("/showExternalProfile", showExternalProfileHandler)
	router.Post("/downloadExternalProfile", downloadExternalProfileHandler)
	router.Post("/aigenerateProfile", aigenerateHandler)

	http.ListenAndServe(":8000", router)
}
