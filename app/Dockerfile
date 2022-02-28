FROM golang:1.17-alpine



ADD . /go/src/app

RUN apk update && apk add --no-cache git jq bash gcc musl-dev build-base alpine-sdk build-base

#RUN mkdir/go/src/app

WORKDIR /go/src/app

#COPY /app /go/src/app
RUN go install github.com/cosmtrek/air@v1.27.3
#    go install github.com/go-delve/delve/cmd/dlv@latest

CMD "sh" "-c" "go mod tidy && air -c .air.toml"
