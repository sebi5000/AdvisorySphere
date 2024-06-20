package model

import (
	"time"
)

type ProjectRequest struct {
	Title           string
	Description     string
	Start           time.Time
	DurationInMonth int
	Vendor          string
	PublishedAt     time.Time
	Industry        string
	Tasks           []string
	Qualifications  []string
}
