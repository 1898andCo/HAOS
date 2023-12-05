FROM golang:1.20.11-alpine3.18 AS gobuild
RUN apk --update add \
    git=2.40.1-r0\
    gcc=12.2.1_git20220924-r10 \
    linux-headers=6.3-r0 \
    musl-dev=1.2.4-r2 \
    make=4.4.1-r1 \
    libseccomp=2.5.4-r2 \
    libseccomp-dev=2.5.4-r2 \
    bash=5.2.15-r5
COPY gobuild /usr/bin/
RUN rm -f /bin/sh && ln -s /bin/bash /bin/sh
WORKDIR /output