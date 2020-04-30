FROM golang:alpine AS builder

RUN apk -U add git
RUN apk -U add bash
RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/jinzhu/gorm
RUN go get -u github.com/gin-contrib/cors

ADD . /go/src/work
WORKDIR /go/src/work

RUN go build -tags netgo -o server main.go

FROM busybox
ENV GOPATH=/usr/local/bin/
COPY --from=builder /go/src/work /usr/local/bin/
RUN chmod +x /usr/local/bin/exec-webapp.sh
ENTRYPOINT ["/usr/local/bin/exec-webapp.sh"]
