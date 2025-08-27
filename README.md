# GoIAChat 🤖💬

Um cliente de linha de comando em Go para interagir com o modelo Gemini AI da Google, proporcionando conversas interativas com IA generativa.

## 🚀 Para que serve

O GoIAChat permite que você:
- Converse naturalmente com a IA Gemini 1.5 Flash
- Obtenha respostas rápidas para perguntas e prompts
- Tenha um histórico de conversa durante a sessão com cache simples
- Experimente a API Gemini de forma interativa

## 📋 Pré-requisitos

- Go 1.18 ou superior
- Uma chave de API do Google Gemini AI
- Git (para clonar o repositório)

## 🔧 Instalação e Configuração

1. **Clone o repositório**:
   ```bash
   git clone <url-do-repositorio>
   cd goiachat
   ```

2. **Instale as dependências**:
   ```bash
   go mod download
   ```

3. **Configure sua chave de API**:
   - Crie um arquivo `.env` na raiz do projeto
   - Adicione sua chave de API:
     ```
     API_KEY=sua_chave_api_aqui
     ```
   - Obtenha uma chave em: [Google AI Studio](https://makersuite.google.com/app/apikey)

## 🏃‍♂️ Como executar

1. **Execute o programa**:
   ```bash
   go run main.go
   ```

2. **Interaja com o chat**:
   - Digite seu prompt e pressione Enter
   - A IA responderá imediatamente
   - Pressione Ctrl+C para sair

3. **Para compilar**:
   ```bash
   go build -o goiachat
   ./goiachat
   ```

## ✨ Características de Implementação

### 🏗️ Arquitetura
- **Cliente otimizado**: Conexão eficiente com a API Gemini
- **Cache simples**: Armazena respostas frequentes para melhor performance
- **Timeout inteligente**: Prevenção de esperas infinitas (30 segundos)

### 🔧 Tecnologias Utilizadas
- **Go 1.18+**: Linguagem de programação eficiente e concorrente
- **Google Generative AI SDK**: Integração oficial com a API Gemini
- **Bufio Scanner**: Leitura eficiente de entrada do usuário

### ⚡ Funcionalidades
- ✅ Conversação interativa em tempo real
- ✅ Cache de respostas para prompts repetidos
- ✅ Tratamento robusto de erros e timeouts
- ✅ Interface de linha de comando intuitiva
- ✅ Configuração simplificada via variáveis de ambiente

### 🛡️ Tratamento de Erros
- Verificação de chave de API
- Timeout para requisições lentas
- Validação de entradas vazias
- Fallback para respostas nulas

## 📝 Estrutura do Projeto

```
goiachat/
├── main.go          # Arquivo principal da aplicação
├── go.mod          # Dependências do Go
├── go.sum          # Checksums das dependências
├── .env            # Arquivo de configuração (criar manualmente)
└── README.md       # Este arquivo
```

## 🔮 Possíveis Melhorias Futuras

- [ ] Suporte a múltiplos modelos Gemini
- [ ] Histórico de conversa persistente
- [ ] Modo de streaming para respostas parciais
- [ ] Comandos especiais (limpar chat, ajuda, etc.)
- [ ] Suporte a prompts multimodais (imagens + texto)

## 📊 Desempenho

O código foi otimizado para:
- Baixo consumo de memória
- Tempo de resposta rápido
- Conexões eficientes com a API
- Caching inteligente de respostas

---

**Nota**: Lembre-se de nunca compartilhar sua chave de API publicamente. O arquivo `.env` está listado no `.gitignore` por padrão para evitar commits acidentais.