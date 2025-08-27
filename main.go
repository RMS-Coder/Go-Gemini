package main

import (
    "bufio"       // Para leitura eficiente de entrada do usuário
    "context"     // Para gerenciamento de contextos e timeouts
    "fmt"         // Para entrada e saída formatada
    "log"         // Para logging de erros
    "os"          // Para acesso a variáveis de ambiente e STDIN
    "strings"     // Para manipulação de strings
    "time"        // Para controle de timeouts

    // SDK oficial do Google para Generative AI
    "github.com/google/generative-ai-go/genai"
    // Para carregar variáveis de ambiente do arquivo .env
    "github.com/joho/godotenv"
    // Para opções de configuração do cliente Google API
    "google.golang.org/api/option"
)

// Estrutura para armazenar o cliente AI e configurações relacionadas
type AIClient struct {
    model  *genai.GenerativeModel  // Modelo generativo da AI
    ctx    context.Context         // Contexto para operações
    cache  map[string]string       // Cache simples para respostas frequentes
}

func main() {
    // Carrega variáveis de ambiente do arquivo .env
    godotenv.Load()
    // Cria um contexto base para a aplicação
    ctx := context.Background()
    
    // Configura e inicializa o cliente da API
    client := setupClient(ctx)
    // Garante que o cliente será fechado ao término do programa
    defer client.Close()

    // Inicializa a estrutura do cliente AI com configurações
    ai := &AIClient{
        model: client.GenerativeModel("gemini-1.5-flash"),  // Usa o modelo Gemini 1.5 Flash
        ctx:   ctx,                                         // Contexto da aplicação
        cache: make(map[string]string),                     // Inicializa o cache vazio
    }
    // Configura a temperatura do modelo para controle de criatividade
    ai.model.SetTemperature(0.7)

    // Mensagem de boas-vindas para o usuário
    fmt.Println("🧠 Bem-vindo ao GoIAChat! Digite seu prompt abaixo. Pressione Ctrl+C para sair.")
    // Inicia o loop principal de chat
    runChatLoop(ai)
}

// setupClient configura e retorna o cliente da API Gemini
func setupClient(ctx context.Context) *genai.Client {
    // Cria um novo cliente com a chave API das variáveis de ambiente
    client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
    // Verifica se houve erro na criação do cliente
    if err != nil {
        log.Fatal("❌ Erro ao criar cliente:", err)
    }
    return client
}

// runChatLoop executa o loop principal de interação com o usuário
func runChatLoop(ai *AIClient) {
    // Cria um scanner para ler entrada do usuário de forma eficiente
    scanner := bufio.NewScanner(os.Stdin)
    
    // Loop infinito para manter a conversação
    for {
        // Prompt para o usuário digitar
        fmt.Print("\n📝 Prompt: ")
        // Lê a entrada do usuário
        if !scanner.Scan() {
            break  // Sai do loop se não houver entrada
        }

        // Remove espaços em branco do início e fim do input
        input := strings.TrimSpace(scanner.Text())
        // Verifica se o input está vazio
        if input == "" {
            fmt.Println("⚠️ Entrada vazia. Tente novamente.")
            continue  // Volta ao início do loop
        }

        // Verifica se a resposta está em cache
        if response, found := ai.cache[input]; found {
            fmt.Printf("🤖 Resposta (cache): %s\n", response)
            continue  // Usa resposta cacheada e volta ao loop
        }

        // Gera resposta da AI para o input
        response, err := ai.generateResponse(input)
        // Verifica se houve erro na geração
        if err != nil {
            fmt.Println("❌ Erro:", err)
            continue  // Volta ao início do loop em caso de erro
        }

        // Verifica se a resposta não está vazia
        if response != "" {
            // Armazena a resposta no cache para futuras consultas
            ai.cache[input] = response
            fmt.Println("🤖 Resposta:", response)
        } else {
            fmt.Println("🤷 Nenhuma resposta gerada.")
        }
    }
}

// generateResponse gera uma resposta para o prompt com timeout
func (ai *AIClient) generateResponse(prompt string) (string, error) {
    // Cria um contexto com timeout de 30 segundos
    ctx, cancel := context.WithTimeout(ai.ctx, 30*time.Second)
    // Garante que a função cancel será chamada ao retornar
    defer cancel()

    // Gera conteúdo usando o modelo AI
    resp, err := ai.model.GenerateContent(ctx, genai.Text(prompt))
    // Verifica se houve erro na geração
    if err != nil {
        return "", err
    }

    // Extrai texto da resposta se disponível
    if resp != nil && len(resp.Candidates) > 0 {
        // Pega o primeiro candidato da resposta
        candidate := resp.Candidates[0]
        // Verifica se o conteúdo existe e tem partes
        if candidate.Content != nil && len(candidate.Content.Parts) > 0 {
            // Converte a parte para texto
            if text, ok := candidate.Content.Parts[0].(genai.Text); ok {
                return string(text), nil
            }
        }
    }
    
    // Retorna string vazia se não houver resposta
    return "", nil
}