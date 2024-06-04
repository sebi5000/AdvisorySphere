package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	router.Get("/", indexHandler)
	router.Get("/showExternalProfile", showExternalProfileHandler)

	router.Post("/project_request", projectRequestHandler)
	router.Post("/downloadExternalProfile", downloadExternalProfileHandler)
	router.Post("/aigenerateProfile", aibeautifyHandler)
	router.Post("/project_clear", clearHandler)

	http.ListenAndServe(":8000", router)
}
