# syntax=docker/dockerfile:1

FROM golang:1.24 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go fmt ./...
RUN go build -o /app/main ./cmd/main.go

FROM gcr.io/distroless/base-debian12 AS final

WORKDIR /root/
COPY --from=builder /app/main /root/main

CMD ["/root/main"]

