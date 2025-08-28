package utils

import (
    "bufio"
    "os"
    "strings"
)

// LoadEnv lê o arquivo .env e carrega as variáveis de ambiente
func LoadEnvFile() error {
    file, err := os.Open(".env")
    if err != nil {
        return err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        // Ignora comentários e linhas vazias
        if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
            continue
        }
        parts := strings.SplitN(line, "=", 2)
        if len(parts) == 2 {
            key := strings.TrimSpace(parts[0])
            value := strings.TrimSpace(parts[1])
            os.Setenv(key, value)
        }
    }
    return scanner.Err()
}

// GetAPIKey retorna o valor da variável de ambiente API_KEY
func GetEnvVariable( key string ) string {
    return os.Getenv(key)
}