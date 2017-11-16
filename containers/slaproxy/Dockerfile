FROM alpine:3.5

RUN apk add --update ca-certificates tzdata && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*

FROM scratch

COPY --from=0 /etc/ssl/certs /etc/ssl/certs
COPY --from=0 /usr/share/zoneinfo /usr/share/zoneinfo

COPY slaproxy /slaproxy

ENTRYPOINT ["/slaproxy"]