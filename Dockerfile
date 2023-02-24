FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy && \
    go mod vendor && \
    mkdir -p out/bin && \
    go build -o ./out/bin/vending_machine_system_control ./cmd/main.go

FROM alpine:3.15

WORKDIR /app

COPY --from=builder /app/out/bin/vending_machine_system_control .
COPY --from=builder /app/env/.env ./env/

EXPOSE 8080

CMD ["./vending_machine_system_control"]