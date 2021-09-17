FROM golang:1.17-alpine3.13 as builder

COPY . /go/src/app

WORKDIR /go/src/app/cmd/app

RUN \
    go version && \
    go install

FROM alpine:3.13

COPY --from=builder /go/bin/app /srv/app

WORKDIR /srv

CMD ["/srv/app"]
