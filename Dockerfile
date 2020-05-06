FROM golang:1.14-alpine

WORKDIR $GOPATH/src/github.com/francisco-serrano/api-golang-docker

ADD ./go.mod $GOPATH/src/github.com/francisco-serrano/api-golang-docker
RUN go get

ADD ./main.go $GOPATH/src/github.com/francisco-serrano/api-golang-docker
RUN go build -o api

EXPOSE 8080

ENTRYPOINT ./api