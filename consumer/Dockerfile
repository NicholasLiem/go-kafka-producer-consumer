# Building Stage
FROM golang:1.18 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64

RUN export GO111MODULE=on

RUN go build -o main .

# Running Stage
FROM alpine:latest

RUN apk add --no-cache musl-dev libc6-compat

WORKDIR /

COPY --from=builder /app/main .

RUN addgroup -g 1000 nonroot && adduser -u 1000 -G nonroot -s /bin/sh -D nonroot

RUN chmod +x main

USER nonroot

CMD ["./main"]