# build ui
FROM ckeyer/dev:vue AS uidist

COPY ui /opt/diego

RUN cd /opt/diego && \
	make build-ui

# build server binary
FROM ckeyer/go:1.10 AS gobin

COPY . /go/src/github.com/ckeyer/diego
COPY --from=uidist /opt/diego/dist /go/src/github.com/ckeyer/diego/ui/dist

RUN cd /go/src/github.com/ckeyer/diego && \
	make build

# ckeyer/diego
FROM alpine:edge

MAINTAINER Chuanjian Wang <me@ckeyer.com>

RUN apk add --update redis ca-certificates

COPY tools/etc/redis.conf /etc/redis.conf
COPY --from=gobin /go/bin/diego /usr/bin/

CMD ["/usr/bin/diego"]
