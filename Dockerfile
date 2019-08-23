FROM ubuntu:18.04

RUN set -x \
    && apt-get update \
    && apt-get install -y \
        mysql-client \
        postgresql-client \
        openssh-client \
        gawk \
        rsync \
        bash

WORKDIR /app

ADD ./*.sh /app/

RUN set -x \
    && chmod +x ./*.sh

ENTRYPOINT ["./backup-start.sh"]
