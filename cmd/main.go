package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	router.Get("/", indexHandler)
	router.Get("/showExternalProfile", showExternalProfileHandler)
	router.Get("/aigenerateProfile", aibeautifyHandler)
	router.Get("/downloadExternalProfile", downloadExternalProfileHandler)

	router.Post("/project_request", projectRequestHandler)
	router.Post("/project_clear", clearHandler)

	http.ListenAndServe(":8000", router)

	//NEXT-STEP: REQUEST Model bauen inkl. Service, welcher mit AI aus einer Projektbeschreibung strukturierte Daten ableitet.
	//Dann darauf das Matching optimieren und entsprechend das Profil optimal anpassen mit AI
}
