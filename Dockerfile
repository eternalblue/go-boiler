FROM golang:1.17-alpine as builder

WORKDIR /app

COPY ./ ./

RUN go mod download && CGO_ENABLED=0 GOOS=linux go build -o api cmd/api/main.go

FROM alpine
ENV GIN_MODE="release"

COPY --from=builder /app/api /app

WORKDIR /app

ENTRYPOINT "./api"