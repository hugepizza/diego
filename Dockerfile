# ckeyer/diego
FROM alpine:edge

MAINTAINER Chuanjian Wang <me@ckeyer.com>

RUN apk add --update redis ca-certificates

ADD tools/etc/redis.conf /etc/redis.conf
ADD bundles/diego /usr/bin/

CMD ["/usr/bin/diego"]
