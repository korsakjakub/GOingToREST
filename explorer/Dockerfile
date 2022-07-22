FROM golang:1.18 as build

WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static-debian11

ENV CONFIG_PATH=/
ENV CONFIG_NAME="test_config"
ENV CONFIG_TYPE="json"

COPY --from=build /go/bin/app /
CMD ["/app"]