FROM golang:1.13
WORKDIR /app
COPY go.mod ./

COPY *.go ./
COPY main.go .

CMD go build && go run main.go