FROM golang:1.17-alpine
RUN apk update
WORKDIR /codingtest/app

COPY . .

EXPOSE 8888

RUN go build -o app main.go

RUN adduser -S golang

USER golang

ENTRYPOINT [ "./app" ]

