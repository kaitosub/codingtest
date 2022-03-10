FROM golang:1.17-alpine

#
# ADD . /go/src/app
#
# RUN apk update
# RUN apk add --no-cache git jq bash gcc musl-dev build-base alpine-sdk build-base
#
# #RUN mkdir/go/src/app
#
# WORKDIR /go/src/app
#
# #COPY /app /go/src/app
# RUN go install github.com/cosmtrek/air@v1.27.3
# RUN go install github.com/go-delve/delve/cmd/dlv@master
#
# CMD "sh" "-c" "go mod tidy && air -c .air.toml"

RUN apk update
WORKDIR /codingtest/app

COPY . .

EXPOSE 8888

RUN go build -o app main.go

RUN adduser -S golang

USER golang

ENTRYPOINT [ "./app" ]

