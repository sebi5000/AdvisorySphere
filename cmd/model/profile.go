package model

type DbProfile struct {
	Id                   string `json:"id"`
	FirstName            string `json:"first_name"`
	LastName             string `json:"last_name"`
	Role                 string `json:"role"`
	Bio                  string `json:"bio"`
	CustomerVoice        string `json:"customer_voice"`
	CustomerReference    string `json:"customer_reference"`
	ProjectTitle         string `json:"project_title"`
	ProjectDescription   string `json:"project_description"`
	ProjectIndustry      string `json:"project_industry"`
	ProjectDurationMonth int    `json:"project_duration"`
	SpecialKnowledge     string `json:"special_knowledge_name"`
	Certificate          string `json:"certificate_name"`
}

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
	Voice     string
	Reference string
}

type Project struct {
	Industry         string
	DurationMonth    int
	Title            string
	ShortDescription string
	Description      string
}
