FROM golang:1.21

WORKDIR /go/src/app

COPY go.mod ./
RUN go mod download && go mod verify

CMD ["go", "run", "main.go"]