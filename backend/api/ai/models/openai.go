package models

import (
	// "fmt"
	"context"
	"enjoy/api/ai/prompts"
	"enjoy/api/ai/config"
	"fmt"
	"log"
	"os"
	"encoding/json"

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

type CommentModResponse struct {
	Allow bool `json:"Allow"`
	Reason string `json:"Reason"`
}

func New(config *config.Config) *OpenAI {
	return &OpenAI{
		Client: openai.NewClient(
			option.WithAPIKey(config.APIKey),
		),
		Model: os.Getenv("OPENAI_MODEL"),
	}
}

func (ai *OpenAI) ModerateMessages(message string) (CommentModResponse, error) {
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
	fmt.Println(res)
	fmt.Println(res.OutputText())

	var stRes CommentModResponse
	json.Unmarshal([]byte(res.OutputText()), &stRes)

	return stRes, nil
}