FROM golang:1.23.2-alpine AS builder

WORKDIR /app

COPY ../../../go.mod ../../../go.sum ./

RUN go mod download && go mod verify

COPY ../../../cmd ./cmd
COPY ../../../data ./data
COPY ../../../internal/auth-service ./internal/auth-service
COPY ../../../pkg ./pkg

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/auth-service/main.go

FROM scratch

WORKDIR /root

COPY --from=builder /app/main .
COPY --from=builder /app/data ./data

ENTRYPOINT ["./main"]