# Primeiro Estágio: Construir a Aplicação
FROM golang:latest AS builder
WORKDIR /
COPY . .
# Compilar a aplicação para um binário estático
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o myapp


# Segundo Estágio: Criar uma Imagem Mínima
FROM scratch
# Copiar o binário compilado do primeiro estágio
COPY --from=builder /app/myapp /myapp
# Definir o ponto de entrada para o binário
ENTRYPOINT [ "./myapp" ]
