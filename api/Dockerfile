FROM golang:1.22.5-alpine AS builder

WORKDIR /src/api

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /src/api/main .

EXPOSE 3000

CMD ["./main"]