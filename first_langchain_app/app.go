package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms"
	_ "github.com/tmc/langchaingo/llms/googleai"
	"github.com/tmc/langchaingo/llms/openai"
)

func main() {

	ctx := context.Background()

	err := godotenv.Load(`first_langchain_app/.env`)
	if err != nil {
		log.Fatal(err)

	}

	llm, err := openai.New(
		openai.WithToken(os.Getenv("GROQ_API_KEY")),
		openai.WithBaseURL("https://api.groq.com/openai/v1"),
		openai.WithModel("llama-3.3-70b-versatile"),
	)

	// llm, err := googleai.New(
	// 	ctx,
	// 	googleai.WithAPIKey(os.Getenv("GOOGLEAI_API_KEY")),
	// 	googleai.WithDefaultModel("gemini-2.5-flash"),
	// )

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
