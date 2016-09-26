FROM golang:1.6

RUN go get github.com/chasinglogic/tessera

ENTRYPOINT /go/bin/tessera
