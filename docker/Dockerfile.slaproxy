FROM alpine:3.6 AS env

RUN apk add --no-cache --update ca-certificates tzdata && \
    update-ca-certificates

FROM instrumentisto/glide AS glide

FROM golang:1.9.2-alpine3.6 AS build

RUN apk add --no-cache --update git

COPY --from=glide /usr/local/bin/glide /usr/local/bin
COPY . /go/src/github.com/lestrrat-go/slack
RUN cd /go/src/github.com/lestrrat-go/slack && \
    glide --yaml docker/slaproxy-glide.yaml install && \
    go build -o /slaproxy cmd/slaproxy/slaproxy.go

FROM alpine:3.6

COPY --from=env /etc/ssl/certs /etc/ssl/certs
COPY --from=env /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /slaproxy /slaproxy
COPY docker/slaproxy-wrapper.sh /slaproxy-wrapper.sh

CMD ["/slaproxy-wrapper.sh"]