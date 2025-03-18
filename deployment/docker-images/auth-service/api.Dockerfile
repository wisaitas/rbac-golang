FROM golang:1.23.2-alpine

WORKDIR /app

COPY ../../../auth-service/go.mod ../../../auth-service/go.sum ./

RUN go mod download && go mod verify

COPY ../../../auth-service/cmd ./cmd
COPY ../../../auth-service/data ./data
COPY ../../../auth-service/internal/auth-service ./internal/auth-service
COPY ../../../auth-service/pkg ./pkg

RUN go mod tidy

RUN go build -o main cmd/main.go

RUN chmod +x main

CMD ["./main"]