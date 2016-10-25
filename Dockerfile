FROM golang:latest

RUN mkdir -p /go/src/github.com/praelatus/backend

CMD [ "go", "run", "/go/src/github.com/praelatus/backend/praelatus.go" ]
