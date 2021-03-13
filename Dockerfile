FROM golang:1.16-alpine3.13 as BUILDER
RUN apk add --no-cache git
WORKDIR ${GOPATH}/src/github.com/guessi/mailgun-tracker
COPY . .
RUN go build -o /go/bin/mailgun-tracker

FROM alpine:3.13
COPY --from=BUILDER /go/bin/mailgun-tracker /opt/
EXPOSE 8080
CMD ["/opt/mailgun-tracker"]
