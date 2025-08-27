package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func main() {
	godotenv.Load()
	continueLoop := true
	reader := bufio.NewReader(os.Stdin)
	for continueLoop {
		fmt.Println("Entre com seu prompt na GoIAChat:")
		var inputText string
		inputText, _ = reader.ReadString('\n')
		ctx := context.Background()
		client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
		if err != nil {
			log.Fatal("Err:", err)
		}
		defer client.Close()

		model := client.GenerativeModel("gemini-1.5-flash")
		resp, err := model.GenerateContent(ctx, genai.Text(inputText))
		if err != nil {
			log.Fatal("Err:", err)
		}

		if resp != nil {
			candidates := resp.Candidates
			if candidates != nil {
				for _, candidate := range candidates {
					content := candidate.Content
					if content != nil {
						text := content.Parts[0]
						log.Println("Gemini:", text)
					}
				}
			} else {
				log.Println("No candidates found")
			}
		} else {
			log.Println("Response is nil")
		}
		var continueInput string
		fmt.Println("Deseja continuar? (s/n)")
		fmt.Scanln(&continueInput)
		if continueInput == "n" {
			continueLoop = false
		}
	}
}
