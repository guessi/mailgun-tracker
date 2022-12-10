# mailgun-tracker

[![GoDoc](https://godoc.org/github.com/guessi/mailgun-tracker?status.svg)](https://godoc.org/github.com/guessi/mailgun-tracker)
[![Go Report Card](https://goreportcard.com/badge/github.com/guessi/mailgun-tracker)](https://goreportcard.com/report/github.com/guessi/mailgun-tracker)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/guessi/mailgun-tracker)](https://github.com/guessi/mailgun-tracker/blob/master/go.mod)

[![Docker Stars](https://img.shields.io/docker/stars/guessi/mailgun-tracker.svg)](https://hub.docker.com/r/guessi/mailgun-tracker/)
[![Docker Pulls](https://img.shields.io/docker/pulls/guessi/mailgun-tracker.svg)](https://hub.docker.com/r/guessi/mailgun-tracker/)
[![Build Status](https://cloud.drone.io/api/badges/guessi/mailgun-tracker/status.svg)](https://cloud.drone.io/guessi/mailgun-tracker)


# Prerequisites

- Docker-CE 20.10+

# Usage

    $ cp config.example.yaml config.yaml

    $ vim config.yaml

    $ docker compose pull # make sure your image is up-to-date

    $ docker compose up [-d]

# FAQ

## What event type support?

currently, mailgun-tracker only support "permanent-failure"

# Reference

- [Docker](https://www.docker.com)
- [Dockerfile Reference](https://docs.docker.com/engine/reference/builder/)

# License

[Apache-2.0](LICENSE)
