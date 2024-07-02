package services

import "sphere/cmd/model"

type ConfigService struct {
}

func (cs ConfigService) Load() (model.Config, error) {

	var config model.Config
	config.Prompts.ProfilePrompts.Bio = ""
	config.Prompts.ProfilePrompts.SpecialKnowledge = ""

	return config, nil
}
