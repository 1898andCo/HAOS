FROM golang:1.16-alpine3.13 AS gobuild
RUN apk --update add \
    git=2.30.2-r0\
    gcc=10.2.1_pre1-r3 \
    linux-headers=5.7.8-r0 \
    musl-dev=1.2.2-r0 \
    make=4.3-r0 \
    libseccomp=2.5.1-r1 \
    libseccomp-dev=2.5.1-r1 \
    bash=5.1.0-r0
COPY gobuild /usr/bin/
RUN rm -f /bin/sh && ln -s /bin/bash /bin/sh
WORKDIR /output