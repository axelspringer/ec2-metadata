FROM alpine:3.6
MAINTAINER Sebastian Döll <sebastian.doell@axelspringer.com>

RUN \
    apk --update add ca-certificates \
	&& rm -rf /var/cache/apk/*

ADD \
    /bin/ec2-metadata_0.0.2_linux_amd64 /bin/ec2-metadata

EXPOSE 80
ENTRYPOINT ["ec2-metadata"]
