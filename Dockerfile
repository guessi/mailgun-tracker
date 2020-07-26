FROM golang:1.14-alpine3.12 as BUILDER
RUN apk add --no-cache git
RUN go get -u github.com/guessi/mailgun-tracker
WORKDIR ${GOPATH}/src/github.com/guessi/mailgun-tracker
RUN go build

FROM alpine:3.12
COPY --from=BUILDER /go/bin/mailgun-tracker /opt/
EXPOSE 8080
CMD ["/opt/mailgun-tracker"]
