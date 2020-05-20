FROM golang:1.14-alpine AS build
LABEL maintainer="angeliski@hotmail.com.br"

ENV APP_DIR /opt/api

COPY . ${APP_DIR}
WORKDIR ${APP_DIR}

RUN apk add build-base

# Remove debug and disable cross compile to create a small binary
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/api *.go

FROM alpine
COPY --from=build /app/api /app/api

# Use an unprivileged user
RUN adduser -D -g '' rdapp
USER rdapp

ENTRYPOINT ["/app/api"]
