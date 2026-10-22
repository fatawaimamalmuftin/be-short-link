FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ShortLink ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/ShortLink ./ShortLink
COPY --from=builder /app/public ./public

EXPOSE 3000

CMD ["./ShortLink"]