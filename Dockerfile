FROM golang:1.22.5-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o stress-test ./cmd/stress-test

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/stress-test .

ENTRYPOINT ["/app/stress-test"]