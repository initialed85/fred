FROM golang:1.23.1 AS builder

WORKDIR /srv/

COPY go.mod /srv/
COPY go.sum /srv/
RUN go get ./...

COPY . /srv/
RUN mkdir -p ./bin
RUN CGO_ENABLED=0 go build -o ./bin/repository-syncer ./cmd/repository_syncer

FROM ubuntu:22.04

ENV TZ Australia/Perth
ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update && apt-get install -y tzdata curl ca-certificates

RUN dpkg-reconfigure -f noninteractive tzdata

COPY --from=builder /srv/bin/repository-syncer /srv/bin/repository-syncer

WORKDIR /srv/

ENV DJANGOLANG_API_ROOT=/api

ENTRYPOINT ["/srv/bin/repository-syncer"]
CMD []
