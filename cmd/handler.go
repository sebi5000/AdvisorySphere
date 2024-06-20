package main

import (
	"fmt"
	"net/http"
	htmx "sphere/cmd/HTMX"
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
	description := r.Form.Get("description")

	_ = description

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

	w.Header().Set("HX-Trigger", "onmatchcompleted")
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

	htmxService := htmx.NewService(w)
	htmxService.AddEvent("onaibeautifycompleted")

	peopleNumber := r.URL.Query().Get("peopleNumber")
	description := r.URL.Query().Get("corr_description")

	var projectService services.ProjectService
	request, err := projectService.CreateProjectRequestFromText(description)

	fmt.Println(request.Title)

	if err != nil {
		status := status.Danger(err.Error())
		eventData := status.GetHXTriggerEvent()
		htmxService.AddEvent(eventData)
		return
	}

	var ps services.ProfileService
	profile := ps.GetProfile(peopleNumber)

	err = ps.AIBeautify(description, &profile)

	if err != nil {
		status := status.Danger(err.Error())
		eventData := status.GetHXTriggerEvent()
		htmxService.AddEvent(eventData)
		return
	}

	templ.Handler(external_profile.ExternalProfile(profile)).ServeHTTP(w, r)
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
