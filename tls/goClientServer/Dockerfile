FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o tls .


FROM alpine:latest

COPY --from=builder /app/tls .

CMD ["/tls", "server"]

