FROM golang:alpine

RUN apk -U add git
RUN apk -U add bash

ADD . /go/src/work
WORKDIR /go/src/work

RUN go build -tags netgo -o server main.go

FROM busybox
ENV GOPATH=/usr/local/bin/
COPY --from=builder /go/src/work/exec-webapp.sh /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/exec-webapp.sh"]
