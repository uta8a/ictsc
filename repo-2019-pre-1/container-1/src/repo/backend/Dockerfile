FROM golang:alpine

RUN apk -U add git
RUN apk -U add bash
RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/jinzhu/gorm

ADD . /go/src/work
WORKDIR /go/src/work
