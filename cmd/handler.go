package main

import (
	"net/http"
	"sphere/cmd/model"
	"sphere/cmd/model/status"
	"sphere/cmd/services"
	"sphere/cmd/views"
	"sphere/cmd/views/components/external_profile"
	"sphere/cmd/views/components/project_request"

	"github.com/a-h/templ"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(views.Index()).ServeHTTP(w, r)
}

//--- PROJECT MATCH FINDING HANDLERS ---

func projectRequestHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//description := r.Form.Get("description")

	match := model.Match{
		model.People{"12345", "Max Mustermann", "", "Senior SAP Consultant", "Hier steht eine Bio"},
		90,
		75,
		"KW: 22",
	}

	match2 := model.Match{
		model.People{"67890", "Sabine Musterfrau", "", "Senior Salesforce Consultant", "Hier steht eine Bio"},
		75,
		90,
		"KW: 24,25",
	}

	var matches = []model.Match{match, match2}
	templ.Handler(project_request.ProjectMatchTable(matches)).ServeHTTP(w, r)
}

func clearHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(project_request.ProjectRequest()).ServeHTTP(w, r)
}

func showExternalProfileHandler(w http.ResponseWriter, r *http.Request) {
	peopleNumber := r.URL.Query().Get("peopleNumber")

	var ps services.ProfileService
	profile := ps.GetProfile(peopleNumber)

	templ.Handler(external_profile.ExternalProfile(profile)).ServeHTTP(w, r)
}

func aibeautifyHandler(w http.ResponseWriter, r *http.Request) {

	peopleNumber := r.URL.Query().Get("peopleNumber")
	description := r.URL.Query().Get("description")

	description = "Wer war Deutschlands erster Bundeskanzler?"
	//TODO: Check Description for eval input or validate, that HTMX Include does HTML Sanitize

	var ps services.ProfileService
	profile := ps.GetProfile(peopleNumber)

	err := ps.AIBeautify(description, &profile)

	if err == nil {
		//templ.Handler(external_profile.ExternalProfile(profile)).ServeHTTP(w, r)
	} else {
		status := status.Danger(err.Error())
		status.SetHXTriggerHeader(w)
	}
}

func downloadExternalProfileHandler(w http.ResponseWriter, r *http.Request) {

	peopleNumber := r.URL.Query().Get("peopleNumber")

	var ps services.ProfileService
	profile := ps.GetProfile(peopleNumber)
	err := ps.Download(profile)

	if err != nil {
		status := status.Danger(err.Error())
		status.SetHXTriggerHeader(w)
	}
}
