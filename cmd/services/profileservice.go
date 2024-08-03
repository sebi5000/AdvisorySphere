package services

import (
	"context"
	"fmt"
	"sphere/cmd/model"
	"sphere/cmd/model/security"
	"sphere/cmd/views/components/external_profile"
	"strings"

	"github.com/dcaraxes/gotenberg-go-client/v8"
	"github.com/supabase-community/supabase-go"
)

type ProfileService struct {
}

func (ps ProfileService) GetProfile(peopleNumber string) model.Profile {

	var profile model.Profile

	client, err := supabase.NewClient(security.GetInstance().SUPABASE_URL, security.GetInstance().SUPABASE_KEY, nil)

	if err != nil {
		fmt.Println(err.Error())
	}

	data, count, err := client.From("consultants").Select("*", "exact", false).Execute()

	if count > 0 {
		_ = data
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	var employee model.People
	employee.Id = "12345"
	employee.Name = "Max Mustermann"
	employee.Role = "SAP Senior Consultant"
	employee.PicturePath = "https://gravatar.com/avatar/0?d=wavatar"
	employee.Bio = `Max Mustermann hat sein Studium der Wirtschaftsinformatik 2015 mit dem Master beendet und berät seit diesem Zeitpunkt
	seine Kunden im Umfeld von Salesforce. Max hat maßgeblich die Digitalisierung im deutschen Mittelstand vorangetrieben
	und in den letzten Jahren zahlreiche Kunden mit Hauptsitz in Deutschland bei internationalen Rollouts beraten und begleitet.`

	profile.People = employee

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

func (ps ProfileService) GetProfilePDF(peopleNumber string) ([]byte, error) {

	var pdfDocument []byte
	gclient := &gotenberg.Client{Hostname: "http://localhost:3000"}

	profile := ps.GetProfile(peopleNumber)
	req, err := ps.profileAsGrotenbergRequest(profile)

	if err != nil {
		return pdfDocument, err
	}

	resp, err := gclient.Post(req)

	if err != nil {
		return pdfDocument, err
	}

	code, err := resp.Body.Read(pdfDocument)
	_ = code

	return pdfDocument, err
}

func (ps ProfileService) AIBeautify(project model.ProjectRequest, profile *model.Profile) error {
	var ai AIService

	basicprompt := `Du bist Vermittler in einer Personalvermittlung und versuchst Profile deiner Klienten, möglichst gemäß den Vorgaben der Projektbeschreibung auszuwählen.
			 		Auch bei kleinen Formulierungen achtest du auf ein Matching in Referenzen, Projekten und Spezialisierungen, sodass der Kandidat für den Kunden sehr gut
					auf die Position passt. `

	preprompt := basicprompt + "Du sollst nun einen Kandidaten mit folgendem Profil vermitteln:" + profile.People.Bio
	preprompt += "Beachte, dass du die Projektbeschreibung mit der nächsten Nachricht erhälst. Deine Antwort sollte nicht mehr als 50 Wörter umfassen."

	answer, err := ai.SendPromptedRequest(preprompt, project.Description)
	profile.People.Bio = answer

	_ = err

	//Special Knowledge

	preprompt = basicprompt + `Liefere eine Liste von Spezialkenntnissen, die besonders gut auf die Projektbeschreibung passen. Trenne die Liste mit einem ;.
							Ein Beispiel könnte wie folgt aussehen: Sales Cloud;Business Analyse und Prozessdesign;Führung internationaler Teams
							Verwende maximal 3 Elemente in deiner Liste und nicht mehr als 40 Zeichen pro Element. Beachte das du die Projektbeschreibung mit der nächsten Nachricht erhälst.`

	answer, err = ai.SendPromptedRequest(preprompt, project.Description)

	_ = err

	specialKnowledge := strings.Split(answer, ";")

	var specialKnowledgeList []model.SpecialKnowledge

	for _, sk := range specialKnowledge {
		knowledge := model.SpecialKnowledge{Name: sk}
		specialKnowledgeList = append(specialKnowledgeList, knowledge)
	}

	profile.SpecialKnowledges = specialKnowledgeList

	//Customer Voice nicht mit KI ermitteln - kann später einfach aus Datenbank kommen

	return err
}

func (ps ProfileService) Download(profile model.Profile) error {

	gclient := &gotenberg.Client{Hostname: "http://localhost:3000"}
	req, err := ps.profileAsGrotenbergRequest(profile)

	if err != nil {
		return err
	}

	err = gclient.Store(req, "/Users/sebastianessling/Downloads/test.pdf")

	return err
}

func (ps ProfileService) profileAsGrotenbergRequest(profile model.Profile) (*gotenberg.HTMLRequest, error) {

	var req *gotenberg.HTMLRequest

	//Build HTML-String with Content from TEMPL and apply stylesheet
	var htmlContent strings.Builder

	external_profile.ExternalProfilePDF(profile).Render(context.TODO(), &htmlContent)
	profilePDF, err := gotenberg.NewDocumentFromString("profile.html", htmlContent.String())

	if err != nil {
		return req, err
	}

	req = gotenberg.NewHTMLRequest(profilePDF)
	req.PaperSize(gotenberg.PaperDimensions{5.625, 10, gotenberg.IN}) //16:9 Ratio like Powerpoint
	req.Landscape(true)
	req.Margins(gotenberg.PageMargins{0.2, 0.2, 0.2, 0.2, gotenberg.IN})
	req.SkipNetworkIdleEvent()

	return req, err
}
