package main

import (
	"bytes"
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

/*func employeeSearchHandler(w http.ResponseWriter, r *http.Request) {
	var employee = rand.IntN(1000)
	templ.Handler(employee_search.EmployeeResult(employee)).ServeHTTP(w, r)
}*/

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

func showExternalProfile(w http.ResponseWriter, r *http.Request) {

	employeeNumber := r.URL.Query().Get("employeeNumber")
	fmt.Println(employeeNumber)

	var ps services.ProfileService
	var profile = ps.GetProfile("12345")

	templ.Handler(external_profile.ExternalProfile(profile)).ServeHTTP(w, r)
}

func downloadExternalProfile(w http.ResponseWriter, r *http.Request) {
	var ps services.ProfileService
	var profile = ps.GetProfile("12345")

	var htmlContent bytes.Buffer
	external_profile.ExternalProfile(profile).Render(r.Context(), &htmlContent)

	w.Header().Set("Content-Disposition", "attachment; filename=profile.pdf")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))

	/*var buffer bytes.Buffer

	//TODO: Create PDF with Gotenberg and create HTML Template to Show PDF
	gclient := &gotenberg.Client{Hostname: "http://localhost:3000"}

	profile, _ := gotenberg.NewDocumentFromString("profile.html", "<html><h1>Wurst!</h1></html>")

	req := gotenberg.NewHTMLRequest(profile)
	req.PaperSize(gotenberg.A4)
	req.Landscape(true)
	req.Margins(gotenberg.NoMargins)
	req.SkipNetworkIdleEvent()
	//resp, err := gclient.Post(req)

	err := gclient.Store(req, "/Users/sebastianessling/Downloads/test.pdf")*/
}
