# Etapa de build
FROM golang:1.23.2-bullseye AS builder

# Definindo o diretório de trabalho
WORKDIR /app

# Cachear as dependências
COPY go.mod go.sum ./
RUN go mod download

# Copiando o código-fonte
COPY . .

# Compilando o binário
RUN CGO_ENABLED=0 GOOS=linux go build -o /server cmd/server/main.go

# Etapa de produção (imagem mínima)
FROM debian:bullseye-slim

# Definindo o diretório de trabalho
WORKDIR /app

# Copiando o binário da etapa de build
COPY --from=builder /server .

# Expondo a porta da aplicação
EXPOSE 8080

# Comando para rodar o binário
CMD ["./server"]
