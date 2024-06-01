package model

type Profile struct {
	People            People
	Certificates      []Certificate
	SpecialKnowledges []SpecialKnowledge
	CustomerVoice     CustomerVoice
	Projects          []Project
}

type People struct {
	Id          string
	Name        string
	PicturePath string
	Role        string
	Bio         string
}

type Certificate struct {
	Name string
}

type SpecialKnowledge struct {
	Name string
}

type CustomerVoice struct {
	Contact string
	Company string
	Voice   string
}

type Project struct {
	Industry         string
	DurationMonth    int
	Title            string
	ShortDescription string
	Description      string
}
