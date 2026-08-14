package ai

import (	
	"enjoy/api/ai/models"
)

type AI struct {
	OpenAI *models.OpenAI
}

func New(openAI *models.OpenAI) *AI {
	return &AI{
		OpenAI: openAI,
	}
}