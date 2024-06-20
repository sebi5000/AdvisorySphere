package status

import (
	"encoding/json"
	"net/http"
)

const (
	SUCCESS = "success"
	INFO    = "info"
	WARNING = "warning"
	DANGER  = "danger"
)

type Status struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}

func New(level string, message string) Status {
	return Status{level, message}
}

func Success(message string) Status {
	return New(SUCCESS, message)
}

func Info(message string) Status {
	return New(INFO, message)
}

func Warning(message string) Status {
	return New(WARNING, message)
}

func Danger(message string) Status {
	return New(DANGER, message)
}

func (s Status) Error() string {
	return s.Message
}

func (s Status) SetHXTriggerHeader(w http.ResponseWriter) error {
	jsonData, err := s.jsonify()
	w.Header().Set("HX-Trigger", jsonData)

	return err
}

func (s Status) GetHXTriggerEvent() string {
	jsonData, err := s.jsonify()
	_ = err
	return jsonData
}

func (s Status) IsError() bool {
	return s.Level == DANGER
}

func (s Status) jsonify() (string, error) {
	s.Message = s.Error()
	eventMap := map[string]Status{}
	eventMap["onstatuschanged"] = s
	jsonData, err := json.Marshal(eventMap)

	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
