package main

import (
	"net/http"
	"sphere/cmd/model"
	"sphere/cmd/views"
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
