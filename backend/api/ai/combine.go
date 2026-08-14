package ai

import (
	"fmt"
	"log"
)

func (ai *AI) ModerateComments(message string) {
	res, err := ai.OpenAI.ModerateMessages(message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("response:", res)

	if res == "yes" {
		fmt.Println("LETS GO")
	} else {
		fmt.Println("BRUH")
	}
}