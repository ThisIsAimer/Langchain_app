package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

func main() {

	err := godotenv.Load(`first_langchain_app/.env`)
	if err != nil {
		log.Fatal(err)

	}

	llm, err := openai.New()
	if err != nil {
		log.Fatal(err)

	}

	prompt := "why is golang so good?"

	completion, err := llms.GenerateFromSinglePrompt(context.Background(), llm, prompt)
	if err != nil {
		log.Fatal(err)

	}

	log.Println(completion)
}
