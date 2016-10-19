FROM golang:1.7

RUN go get github.com/praelatus/backend

ENTRYPOINT /go/bin/praelatus-server
