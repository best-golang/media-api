FROM golang:1.10.3

RUN apt-get update -qq && \
  apt-get -q -y install mysql-client && \
  go get -u github.com/golang/dep/cmd/dep

RUN mkdir -p /go/src/api
ADD . /go/src/api
RUN rm -rf /go/src/api/.vendor-new

WORKDIR /go/src/api
