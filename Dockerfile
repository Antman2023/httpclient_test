# dedicatedproxy on docker
FROM golang:1.5.1
MAINTAINER pathletboy <pathletboy@gmail.com>

COPY main.go $GOPATH/src/app/
WORKDIR $GOPATH/src/app/

EXPOSE 80
EXPOSE 8888

RUN go get && go install
ENTRYPOINT $GOPATH/bin/app