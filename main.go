package main

import (
    "bufio"
    "context"
    "fmt"
    "log"
    "os"
    "strings"

    "github.com/google/generative-ai-go/genai"
    "github.com/joho/godotenv"
    "google.golang.org/api/option"
)

func main() {
    godotenv.Load()

    ctx := context.Background()
    client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
    if err != nil {
        log.Fatal("Erro ao criar cliente:", err)
    }
    defer client.Close()

    model := client.GenerativeModel("gemini-1.5-flash")
    reader := bufio.NewReader(os.Stdin)

    fmt.Println("🧠 Bem-vindo ao GoIAChat! Digite seu prompt abaixo. Pressione Ctrl+C para sair.")

    for {
        fmt.Print("\n📝 Prompt: ")
        inputText, _ := reader.ReadString('\n')
        inputText = strings.TrimSpace(inputText)

        if inputText == "" {
            fmt.Println("⚠️ Entrada vazia. Tente novamente.")
            continue
        }

        resp, err := model.GenerateContent(ctx, genai.Text(inputText))
        if err != nil {
            fmt.Println("❌ Erro ao gerar resposta:", err)
            continue
        }

        if resp != nil && resp.Candidates != nil {
            for _, candidate := range resp.Candidates {
                if candidate.Content != nil {
                    fmt.Println("🤖 Resposta:", candidate.Content.Parts[0])
                }
            }
        } else {
            fmt.Println("🤷 Nenhuma resposta gerada.")
        }
    }
}
