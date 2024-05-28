package services

import (
	"context"
	"errors"
	"sphere/cmd/model"
	"sphere/cmd/views/components/external_profile"
	"strings"

	"github.com/dcaraxes/gotenberg-go-client/v8"
)

type ProfileService struct {
}

func (ps ProfileService) GetProfile(employeeNumber string) model.Profile {

	var profile model.Profile

	var employee model.Employee
	employee.Id = "12345"
	employee.Name = "Max Mustermann"
	employee.Role = "SAP Senior Consultant"
	employee.PicturePath = "https://source.unsplash.com/schwarzer-hund-tragt-blaues-jeanshalsband-K4mSJ7kc0As"
	employee.Bio = `Max Mustermann hat sein Studium der Wirtschaftsinformatik 2015 mit dem Master beendet und berät seit diesem Zeitpunkt
	seine Kunden im Umfeld von Salesforce. Max hat maßgeblich die Digitalisierung im deutschen Mittelstand vorangetrieben
	und in den letzten Jahren zahlreiche Kunden mit Hauptsitz in Deutschland bei internationalen Rollouts beraten und begleitet.`

	profile.Employee = employee

	var customerVoice model.CustomerVoice
	customerVoice.Company = "Musterfirma GmbH"
	customerVoice.Contact = "Rene Rakete"
	customerVoice.Voice = "Max hat uns mit seiner Performance wirklich vorangebracht. Neben bester Qualität, hat es vor allem menschlich sofort gepasst."

	profile.CustomerVoice = customerVoice

	var cert model.Certificate
	cert.Name = "Salesforce Administrator"

	var cert2 model.Certificate
	cert2.Name = "Salesforce Sales Cloud Consultant"

	var cert3 model.Certificate
	cert3.Name = "Salesforce AI Practioner"

	profile.Certificates = append(profile.Certificates, cert, cert2, cert3)

	var sk model.SpecialKnowledge
	sk.Name = "Sales Cloud"

	var sk2 model.SpecialKnowledge
	sk2.Name = "Business Analyse und Prozessdesign"

	profile.SpecialKnowledges = append(profile.SpecialKnowledges, sk, sk2)

	var project model.Project
	project.Industry = "Chemieindustrie"
	project.DurationMonth = 12
	project.Title = "Einführung Salesforce Sales Cloud"
	project.ShortDescription = "Hier steht eine Beschreibung des Projektinhaltes, welche durchaus mehrere Zeilen haben kann und die Tätigkeiten während des Projektes beschreibt."
	project.Description = "Hier steht eine Beschreibung des Projektinhaltes, welche durchaus mehrere Zeilen haben kann und die Tätigkeiten während des Projektes beschreibt."

	var project2 model.Project
	project2.Industry = "Retail"
	project2.DurationMonth = 18
	project2.Title = "Integration Salesforce & SAP"
	project2.ShortDescription = "Hier steht eine Beschreibung des Projektinhaltes, welche durchaus mehrere Zeilen haben kann und die Tätigkeiten während des Projektes beschreibt."
	project2.Description = "Hier steht eine Beschreibung des Projektinhaltes, welche durchaus mehrere Zeilen haben kann und die Tätigkeiten während des Projektes beschreibt."

	profile.Projects = append(profile.Projects, project, project2)

	return profile
}

func (ps ProfileService) AIGenerate(projectDescription string, profile *model.Profile) error {

	return errors.New("TODO: Implement Method")
}

func (ps ProfileService) Download(profile model.Profile) error {

	//Build HTML-String with Content from TEMPL and apply stylesheet
	var htmlContent strings.Builder
	htmlContent.WriteString(`<html><head>
									 <link rel="stylesheet" href="picostyle.css"/>
									 <link rel="stylesheet" href="flexboxstyle.css"/>
								 </head><body>`)

	external_profile.ExternalProfile(profile).Render(context.TODO(), &htmlContent)
	htmlContent.WriteString(`</body></html>`)

	gclient := &gotenberg.Client{Hostname: "http://localhost:3000"}
	profilePDF, err := gotenberg.NewDocumentFromString("profile.html", htmlContent.String())

	if err != nil {
		return err
	}

	picoStyle, err := gotenberg.NewDocumentFromPath("picostyle.css", "/Users/sebastianessling/Documents/AdvisorySphere/assets/pico.min.css")

	if err != nil {
		return err
	}

	flexboxStyle, err := gotenberg.NewDocumentFromPath("flexboxstyle.css", "/Users/sebastianessling/Documents/AdvisorySphere/assets/flexboxgrid.min.css")

	if err != nil {
		return err
	}

	req := gotenberg.NewHTMLRequest(profilePDF)
	req.Assets(picoStyle, flexboxStyle)
	req.PaperSize(gotenberg.A3)
	req.Landscape(true)
	req.Margins(gotenberg.NoMargins)
	req.SkipNetworkIdleEvent()

	err = gclient.Store(req, "/Users/sebastianessling/Downloads/test.pdf")

	return err
}
