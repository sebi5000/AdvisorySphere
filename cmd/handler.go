package main

import (
	"fmt"
	"net/http"
	"sphere/cmd/model"
	"sphere/cmd/services"
	"sphere/cmd/views"
	"sphere/cmd/views/components/external_profile"
	"sphere/cmd/views/components/project_request"

	"github.com/a-h/templ"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(views.Index()).ServeHTTP(w, r)
}

func projectRequestHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//description := r.Form.Get("description")

	match := model.Match{
		"Max Mustermann",
		90,
		75,
		"KW: 22",
	}

	match2 := model.Match{
		"Sabine Musterfrau",
		75,
		90,
		"KW: 24,25",
	}

	var matches = []model.Match{match, match2}
	templ.Handler(project_request.ProjectMatchTable(matches)).ServeHTTP(w, r)
}

func showExternalProfileHandler(w http.ResponseWriter, r *http.Request) {

	employeeNumber := r.URL.Query().Get("employeeNumber")

	var ps services.ProfileService
	var profile = ps.GetProfile(employeeNumber)

	templ.Handler(external_profile.ExternalProfile(profile)).ServeHTTP(w, r)
}

func aigenerateHandler(w http.ResponseWriter, r *http.Request) {

	employeeNumber := r.URL.Query().Get("employeeNumber")

	var ps services.ProfileService
	profile := ps.GetProfile(employeeNumber)
	description := "HIER EINE BEISPIELHAFTE PROJEKTBESCHREIBUNG EINFÜGEN, BIS DIE VOM FRONTEND ÜBERNOMMEN WIRD."
	ps.AIGenerate(description, &profile)

	templ.Handler(external_profile.ExternalProfile(profile)).ServeHTTP(w, r)
}

func downloadExternalProfileHandler(w http.ResponseWriter, r *http.Request) {

	employeeNumber := r.URL.Query().Get("employeeNumber")

	var ps services.ProfileService
	profile := ps.GetProfile(employeeNumber)
	err := ps.Download(profile)

	if err != nil {
		fmt.Println(err.Error())
	}
}
