FROM golang:1.27.1 AS builder

WORKDIR /app

COPY go.sum go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./ShortLink ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/be-golang ./ShortLink

COPY --from=builder /app/public ./public

CMD ["/app/ShortLink"]