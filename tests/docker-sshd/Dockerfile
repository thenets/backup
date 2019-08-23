FROM ubuntu:16.04

RUN set -x \
    && apt-get update && apt-get install -y openssh-server \
    && mkdir /var/run/sshd \
    && mkdir -p /root/.ssh \
    && sed -i "s/PermitRootLogin prohibit-password/PermitRootLogin without-password/g" /etc/ssh/sshd_config \
    && echo "AuthorizedKeysFile     %h/.ssh/authorized_keys" >> /etc/ssh/sshd_config

COPY ./id_rsa* /root/.ssh/
COPY ./client.pub /root/.ssh/authorized_keys

EXPOSE 22

CMD ["/usr/sbin/sshd", "-D"]
