FROM golang:1.18 as build

ENV APP poster
ENV CMD_PATH post_producer.go

COPY . $GOPATH/src/$APP
WORKDIR $GOPATH/src/$APP

RUN go build -v -o /$APP $GOPATH/src/$APP/$CMD_PATH

FROM alpine:3.14
ENV APP poster
COPY --from=build /$APP_NAME .

CMD ./$APP