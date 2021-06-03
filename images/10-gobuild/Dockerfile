FROM golang:1.16-alpine3.12 AS gobuild
RUN apk --update -U add \
    git=2.26.3-r0 \
    gcc=9.3.0-r2 \
    linux-headers=5.4.5-r1 \
    musl-dev=1.1.24-r10 \
    make=4.3-r0 \
    libseccomp=2.4.4-r0 \
    libseccomp-dev=2.4.4-r0 \
    bash=5.0.17-r0
COPY gobuild /usr/bin/
RUN rm -f /bin/sh && ln -s /bin/bash /bin/sh
WORKDIR /output