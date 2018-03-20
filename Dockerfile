# ckeyer/diego
FROM alpine:edge

MAINTAINER Chuanjian Wang <me@ckeyer.com>

RUN apk add --update redis ca-certificates

ADD bundles/diego /usr/bin/

CMD ["/usr/bin/diego"]
