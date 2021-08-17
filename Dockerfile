#Builder Stage
FROM golang:alpine3.14 AS builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o twitch-lichess-predictions

#Production Stage
FROM alpine:3.14.1 AS production

COPY --from=builder /app/twitch-lichess-predictions .

ENTRYPOINT ["./twitch-lichess-predictions"]