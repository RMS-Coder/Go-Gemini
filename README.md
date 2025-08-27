
# Utilizando-gemini-com-go

Aplicação com o intuito de conectar o google Gemini utilizando a linguagem Go

  

## Requisitos

  

- Go 1.20+ (ou versão compatível)

- Conta no Google Cloud com acesso à API do Google Generative AI

- API Key configurada no arquivo `.env`

- Dependências do projeto:

- [google/generative-ai-go](https://github.com/google/generative-ai-go)

- [joho/godotenv](https://github.com/joho/godotenv)

  

## Instalação

  

1. Clone o repositório:

  

```bash

git clone https://github.com/GabouKing/utilizando-gemini-com-go

cd utilizando-gemini-com-go
```
  

2. Instale as dependências:

```bash

go mod tidy
```
  

3. Crie um arquivo .env na raiz do projeto e adicione a sua chave de API do Google::

```bash

API_KEY=your_api_key_here
```
  

## Uso

Para executar a aplicação, utilize o comando:

  

```bash
go  run  main.go
```
O  programa  solicitará  que  você  insira  um  prompt  no  terminal.  Após  digitar  o  texto  e  pressionar  "Enter",  ele  enviará  o  prompt  para  a  API  do  modelo  gemini-1.5-flash  e  exibirá  a  resposta  gerada.

## Exemplo de Uso
```bash
Entre com seu prompt na GoIAChat:
Qual é a capital da França?
Gemini: A capital da França é Paris.

Deseja continuar? (s/n)

```
Digite "s" para continuar inserindo novos prompts ou "n" para encerrar o programa.