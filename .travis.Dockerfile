FROM ubuntu:18.04

RUN apt-get -qq update && \
    apt-get install -y git make btrfs-tools libdevmapper-dev libgpgme-dev libostree-dev

ADD https://storage.googleapis.com/golang/go1.11.12.linux-amd64.tar.gz /tmp

RUN tar -C /usr/local -xzf /tmp/go1.11.12.linux-amd64.tar.gz && \
    rm /tmp/go1.11.12.linux-amd64.tar.gz

ENV PATH="/usr/local/go/bin:/root/go/bin:${PATH}"
