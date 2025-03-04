FROM golang:1.21 as builder

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o main ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]