package services

import (
	"context"
	"encoding/json"
	"sphere/cmd/model"
	"sphere/cmd/model/security"
	"sphere/cmd/views/components/external_profile"
	"strings"

	"github.com/dcaraxes/gotenberg-go-client/v8"
	"github.com/supabase-community/supabase-go"
)

type ProfileService struct {
}

func (ps ProfileService) GetProfile(peopleNumber string) (model.Profile, error) {

	var profile model.Profile
	var dbProfile []model.DbProfile

	client, err := supabase.NewClient(security.GetInstance().SUPABASE_URL, security.GetInstance().SUPABASE_KEY, nil)

	if err != nil {
		return profile, err
	}

	//Read Consultant Profile from Database
	//TODO: Filter based on peopleNumber
	data, count, err := client.From("consultant_profile").Select("*", "exact", false).Execute()

	//Check if data exists
	if count > 0 {
		if err == nil {
			err = json.Unmarshal(data, &dbProfile)
		}
	}

	//If error occurs while access database return error
	if err != nil {
		return profile, err
	}

	for i, profileRecord := range dbProfile {

		//Basis consultant information are same for every db record. We only need to fetch them once.
		if i == 1 {
			//Add basis consultant information
			var consultant model.People
			consultant.Id = profileRecord.Id
			consultant.Name = profileRecord.FirstName + ` ` + profileRecord.LastName
			consultant.Role = profileRecord.Role
			consultant.PicturePath = "https://gravatar.com/avatar/0?d=wavatar"
			consultant.Bio = profileRecord.Bio

			profile.People = consultant

			//Add customer voice
			var customerVoice model.CustomerVoice
			customerVoice.Reference = profileRecord.CustomerReference
			customerVoice.Voice = profileRecord.CustomerVoice
			profile.CustomerVoice = customerVoice
		}

		//Add special knowledge
		var specialKnowledge model.SpecialKnowledge
		specialKnowledge.Name = profileRecord.SpecialKnowledge
		ps.addSpecialKnowledgeUnique(specialKnowledge, &profile.SpecialKnowledges)

		//Add projects
		var project model.Project
		project.Industry = profileRecord.ProjectIndustry
		project.DurationMonth = profileRecord.ProjectDurationMonth
		project.Title = profileRecord.ProjectTitle
		project.ShortDescription = profileRecord.ProjectDescription
		project.Description = profileRecord.ProjectDescription
		ps.addProjectUnique(project, &profile.Projects)

		//Add certifications
		var cert model.Certificate
		cert.Name = profileRecord.Certificate
		ps.addCertificateUnique(cert, &profile.Certificates)
	}

	return profile, err
}

func (ps ProfileService) GetProfilePDF(peopleNumber string) ([]byte, error) {

	var pdfDocument []byte
	gclient := &gotenberg.Client{Hostname: "http://localhost:3000"}

	profile, err := ps.GetProfile(peopleNumber)

	if err != nil {
		return pdfDocument, err
	}

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

func (ps ProfileService) addProjectUnique(project model.Project, projects *[]model.Project) {

	var exists = false

	for _, existingProject := range *projects {
		exists = existingProject.Title == project.Title
		if exists {
			break
		}
	}

	if !exists {
		*projects = append(*projects, project)
	}
}

func (ps ProfileService) addSpecialKnowledgeUnique(sk model.SpecialKnowledge, specials *[]model.SpecialKnowledge) {

	var exists = false

	for _, existingSpecial := range *specials {

		exists = existingSpecial.Name == sk.Name

		if exists {
			break
		}
	}

	if !exists {
		*specials = append(*specials, sk)
	}
}

func (ps ProfileService) addCertificateUnique(cert model.Certificate, certificates *[]model.Certificate) {

	var exists = false

	for _, existingCertificate := range *certificates {

		exists = existingCertificate.Name == cert.Name

		if exists {
			break
		}
	}

	if !exists {
		*certificates = append(*certificates, cert)
	}
}
