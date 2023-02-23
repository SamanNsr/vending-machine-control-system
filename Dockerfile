FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o myapp ./cmd/main.go

FROM alpine:3.15

WORKDIR /app

COPY --from=builder /app/myapp .

EXPOSE 8080

CMD ["./myapp"]