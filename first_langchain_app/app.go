package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
)

func main() {

	ctx := context.Background()

	err := godotenv.Load(`first_langchain_app/.env`)
	if err != nil {
		log.Fatal(err)

	}
	key := os.Getenv("GOOGLEAI_API_KEY")
	// llm, err := openai.New()
	llm, err := googleai.New(
		ctx,
		googleai.WithAPIKey(key),
		googleai.WithDefaultModel("gemini-2.5-flash"),
	)

	if err != nil {
		log.Fatal(err)

	}

	prompt := "whats the capital of india?"

	completion, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
	if err != nil {
		log.Fatal(err)

	}

	log.Println(completion)
}
