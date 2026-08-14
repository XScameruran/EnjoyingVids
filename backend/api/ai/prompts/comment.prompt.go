package prompts

import (
	"log"
	"os"
)

func GetComPrompt(message string) string {
	prompt, err := os.ReadFile("api/ai/prompts/comment.txt")

	if err != nil {
		log.Fatal(err)
	}


	return string(prompt) + message
}