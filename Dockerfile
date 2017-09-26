# tibco/mashling
FROM golang:1.9.0-alpine3.6
MAINTAINER Jeffrey Bozek, jbozek@tibco.com

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN go get github.com/constabulary/gb/...
ADD . $GOPATH/src/github.com/TIBCOSoftware/mashling
RUN go install github.com/TIBCOSoftware/mashling/cli/...

RUN mkdir /builds/
WORKDIR /builds/
