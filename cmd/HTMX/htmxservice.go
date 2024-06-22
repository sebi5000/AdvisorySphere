package htmx

import (
	"encoding/json"
	"net/http"
)

type HTMXService struct {
	ResponseWriter http.ResponseWriter
	eventStack     []Event
}

type Event struct {
	Name  string
	Param any
}

func NewService(w http.ResponseWriter) *HTMXService {
	service := &HTMXService{
		ResponseWriter: w,
	}

	return service
}

func (hs *HTMXService) AddEvent(event Event) {

	hs.eventStack = append(hs.eventStack, event)

	events, err := hs.jsonify()

	_ = err

	hs.ResponseWriter.Header().Del("HX-Trigger")
	hs.ResponseWriter.Header().Set("HX-Trigger", events)
}

func (s HTMXService) jsonify() (string, error) {

	eventMap := map[string]any{}

	for _, event := range s.eventStack {
		eventMap[event.Name] = event.Param
	}

	jsonData, err := json.Marshal(eventMap)

	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
