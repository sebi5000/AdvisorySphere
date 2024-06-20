package htmx

import (
	"net/http"
	"strings"
)

type HTMXService struct {
	ResponseWriter http.ResponseWriter
	eventStack     []string
	eventChain     string
}

func NewService(w http.ResponseWriter) *HTMXService {
	service := &HTMXService{
		ResponseWriter: w,
	}

	return service
}

func (hs *HTMXService) AddEvent(event string) {

	hs.eventStack = append(hs.eventStack, event)
	var eventChain string = ""

	for i := range hs.eventStack {
		eventChain += hs.eventStack[i] + ","
	}

	eventChain, found := strings.CutSuffix(eventChain, ",")

	hs.eventChain = eventChain

	_ = found //If not found - no problem

	hs.ResponseWriter.Header().Del("HX-Trigger")
	hs.ResponseWriter.Header().Set("HX-Trigger", hs.eventChain)
}
