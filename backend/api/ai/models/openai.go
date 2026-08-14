package models

import (
	// "fmt"
	"context"
	"enjoy/api/ai/prompts"
	"enjoy/api/ai/config"
	"fmt"
	"log"
	"os"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/responses"
)

type OpenAI struct {
	Client openai.Client
	Model string
}

type Config struct {
	APIKey string
	Model string
}

func New(config *config.Config) *OpenAI {
	return &OpenAI{
		Client: openai.NewClient(
			option.WithAPIKey(config.APIKey),
		),
		Model: os.Getenv("OPENAI_MODEL"),
	}
}

func (ai *OpenAI) ModerateMessages(message string) (string, error) {
	ctx := context.Background()
	messageprompt := prompts.GetComPrompt(message)

	fmt.Println("msg prompt:", messageprompt)
	params := responses.ResponseNewParams{
		Input: responses.ResponseNewParamsInputUnion{
			OfString: openai.String(messageprompt),
		},
		Model: openai.ChatModel(ai.Model),
	}
	res, err := ai.Client.Responses.New(ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	return res.OutputText(), nil
}