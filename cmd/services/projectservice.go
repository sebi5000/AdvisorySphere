package services

import (
	"sphere/cmd/model"
	"strconv"
	"strings"
	"time"
)

type ProjectService struct {
}

func (ps ProjectService) CreateProjectRequestFromText(text string) (model.ProjectRequest, error) {

	var ai AIService

	//Schema, which is used by LLM as return format
	schema := ` START TEMPLATE
				#[TITLE];#[DESCRIPTION];#[DURATION];#[INDUSTRY];#[VENDOR];#[PUBLISHEDAT];#[START];
				END TEMPLATE
			
				In der nachfolgenden Nachricht, befindet sich eine Projektbeschreibung. Extrahiere verschiedene Bestandteile. 
				Die Bestandteile ersetzt du dann mit den Platzhaltern im Template. Anbei die Zuordnung von Platzhaltern und Bestandteilen:
				[TITLE] = Der Titel der Projektbeschreibung
				[DESCRIPTION] = Der beschreibende Text des Projektes
				[DURATION] = Die Dauer des Projekteinsatzes in Monaten. Ist diese nicht direkt angegeben, berechne Sie anhand von Start und Endedatum (Beispiel: 01.01.2024 - 01.07.2024 = 6 Monate). Nimm ausschließlich die Zahl in deine Antwort z.B. 6
				[INDUSTRY] = Die Branche in der das Unternehmen tätig ist. Sollte dies nicht im Text enthalten sein, prüfe ob du es vom Unternehmen ableiten kannst.
				[VENDOR] = Das Unternehmen, welches direkt oder als Vertreter eines Dritten für das Projekt jemanden sucht
				[PUBLISHEDAT] = Das Datum an dem die Projektbeschreibung veröffentlicht wurde
				[START] = Das Startdatum des Einsatzes. Sollte asap angegeben sein, nimm das heutige aktuelle Datum.

				Kannst du ein Bestandteil nicht ermitteln, ersetze den Platzhalter mit "NOT FOUND". Entferne START TEMPLATE und END TEMPLATE in deiner Antwort.`

	answer, err := ai.SendPromptedRequest(schema, text)

	splittedAnswer := strings.Split(answer, ";")

	var projectRequest model.ProjectRequest

	for i, text := range splittedAnswer {
		switch i {
		case 0:
			projectRequest.Title = text
		case 1:
			projectRequest.Description = text
		case 2:
			if text != "NOT FOUND" {
				projectRequest.DurationInMonth, err = strconv.Atoi(text)
			}
		case 3:
			projectRequest.Industry = text
		case 4:
			projectRequest.Vendor = text
		}
	}

	location := time.Now().Location()
	projectRequest.PublishedAt = time.Date(2024, 6, 19, 0, 0, 0, 0, location)
	projectRequest.Start = time.Date(2024, 6, 19, 0, 0, 0, 0, location)
	projectRequest.Tasks = []string{"Durchführung von Payroll Aktivitäten", "Beratung zu Payroll Fragen"}
	projectRequest.Qualifications = []string{"Vertraut mit Payrollanforderungen", "Expertise in SAP HCM"}

	return projectRequest, err
}
