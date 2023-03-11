FROM public.ecr.aws/docker/library/golang:1.20-alpine3.17 as BUILDER
RUN apk add --no-cache git ca-certificates
WORKDIR ${GOPATH}/src/github.com/guessi/mailgun-tracker
COPY . .
RUN GOPROXY=direct CGO_ENABLED=0 go build -o /go/bin/mailgun-tracker

FROM scratch
COPY --from=BUILDER /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=BUILDER /go/bin/mailgun-tracker /opt/
EXPOSE 8080
CMD ["/opt/mailgun-tracker"]
