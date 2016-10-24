FROM ubuntu:latest

RUN apt-get update && apt-get -y upgrade
RUN apt-get install -y python-pip watchdog && pip install https://github.com/joh/when-changed/archive/master.zip
RUN apt-get install -y golang git

ENV GOPATH /go

RUN mkdir -p /go/{src,bin,pkg}
RUN useradd gopher -d /go
RUN go get github.com/praelatus/backend

EXPOSE 8080

USER gopher
WORKDIR /go/src/github.com/chasinglogic/praelatus/backend
ENTRYPOINT when-changed -r . -c go run cmd/praelatus-server/main.go
