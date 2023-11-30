FROM golang:1.17.11-alpine3.16 AS gobuild
RUN apk --update add \
    git=2.36.6-r0\
    gcc=11.2.1_git20220219-r2 \
    linux-headers=5.16.7-r1 \
    musl-dev=1.2.3-r3 \
    make=4.3-r0 \
    libseccomp=2.5.2-r1 \
    libseccomp-dev=2.5.2-r1 \
    bash=5.1.16-r2
COPY gobuild /usr/bin/
RUN rm -f /bin/sh && ln -s /bin/bash /bin/sh
WORKDIR /output