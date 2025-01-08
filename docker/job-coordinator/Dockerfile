FROM golang:1.23.1 AS builder

WORKDIR /srv/

COPY go.mod /srv/
COPY go.sum /srv/
RUN go get ./...

COPY . /srv/
RUN mkdir -p ./bin
RUN CGO_ENABLED=0 go build -o ./bin/job-coordinator ./cmd/job_coordinator

FROM ubuntu:22.04

ENV TZ Australia/Perth
ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update && apt-get install -y tzdata curl ca-certificates

RUN dpkg-reconfigure -f noninteractive tzdata

COPY --from=builder /srv/bin/job-coordinator /srv/bin/job-coordinator

WORKDIR /srv/

ENV DJANGOLANG_API_ROOT=/api

ENTRYPOINT ["/srv/bin/job-coordinator"]
CMD []
