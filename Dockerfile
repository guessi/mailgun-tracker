FROM golang:1.16-alpine3.13 as BUILDER
RUN apk add --no-cache git ca-certificates
WORKDIR ${GOPATH}/src/github.com/guessi/mailgun-tracker
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/mailgun-tracker

FROM scratch
COPY --from=BUILDER /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=BUILDER /go/bin/mailgun-tracker /opt/
EXPOSE 8080
CMD ["/opt/mailgun-tracker"]
