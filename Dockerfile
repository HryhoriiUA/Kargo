FROM golang:1.20-alpine

WORKDIR /app

COPY . .
RUN go mod init kargo
RUN go mod tidy
RUN go build -o main .

ENV PORT=8080

EXPOSE 8080

# Команда для запуска приложения
CMD ["./main"]