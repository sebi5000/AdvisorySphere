package model

type Config struct {
	Prompts struct {
		ProfilePrompts struct {
			Bio              string `json:"bio"`
			SpecialKnowledge string `json:"specialKnowledge"`
		}
	}
}
