# mailgun-tracker

[![Docker Stars](https://img.shields.io/docker/stars/guessi/mailgun-tracker.svg)](https://hub.docker.com/r/guessi/mailgun-tracker/)
[![Docker Pulls](https://img.shields.io/docker/pulls/guessi/mailgun-tracker.svg)](https://hub.docker.com/r/guessi/mailgun-tracker/)
[![Docker Automated](https://img.shields.io/docker/automated/guessi/mailgun-tracker.svg)](https://hub.docker.com/r/guessi/mailgun-tracker/)
[![Build Status](https://cloud.drone.io/api/badges/guessi/mailgun-tracker/status.svg)](https://cloud.drone.io/guessi/mailgun-tracker)


# Prerequisites

- Docker-CE 19.03+
- Docker Compose 1.24.0+

# Usage

    $ cp config.example.yaml config.yaml

    $ vim config.yaml

    $ docker-compose pull # make sure your image is up-to-date

    $ docker-compose up [-d]

# FAQ

## What event type support?

currently, mailgun-tracker only support "permanent-failure"

# Reference

- [Docker CE](https://www.docker.com/community-edition)
- [Docker Compose](https://docs.docker.com/compose/overview/)
- [Dockerfile Reference](https://docs.docker.com/engine/reference/builder/)

# License

[Apache-2.0](LICENSE)
