package status

import (
	"encoding/json"
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

func (s Status) IsError() bool {
	return s.Level == DANGER
}

func (s Status) ToJSON() string {
	json, err := s.jsonify()
	_ = err
	return json
}

func (s Status) jsonify() (string, error) {
	jsonData, err := json.Marshal(s)

	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
