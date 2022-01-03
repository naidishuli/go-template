FROM golang:1.17

ADD . /go-template
WORKDIR /go-template

RUN go build -o cmd/main cmd/main.go

EXPOSE 7001
CMD ["./cmd/main"]