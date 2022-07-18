FROM golang:1.18 as build

ENV APP saver
ENV CMD_PATH saver.go

COPY . $GOPATH/src/$APP
WORKDIR $GOPATH/src/$APP

RUN go build -v -o /$APP $GOPATH/src/$APP/$CMD_PATH

FROM alpine:3.14
ENV APP saver
COPY --from=build /$APP_NAME .

CMD ./$APP