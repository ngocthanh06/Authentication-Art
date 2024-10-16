# Stage 1: Build the application
FROM golang:1.23-alpine as builder
RUN apk update && apk upgrade && \
    apk --update add git make bash build-base

RUN apk add --no-cache git build-base

ENV DB_CONN=postgres
ENV DB_HOST=db
ENV DB_NAME=authentication
ENV DB_PASS=pass
ENV DB_USER=user
ENV PORT=5432
ENV APP_PORT=8080

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ../ .

RUN go install github.com/air-verse/air@latest

RUN go build -o main ./app/authentication

CMD ["air"]
