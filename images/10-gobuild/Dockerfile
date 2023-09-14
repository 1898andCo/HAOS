FROM golang:1.17.0-alpine
RUN apk --update add \
    git=2.32.7-r0 \
    gcc=10.3.1_git20210424-r2 \
    linux-headers=5.10.41-r0 \
    musl-dev=1.2.2-r5 \
    make=4.3-r0 \
    libseccomp=2.5.1-r2 \
    libseccomp-dev=2.5.1-r2 \
    bash=5.1.16-r0 
COPY gobuild /usr/bin/
RUN rm -f /bin/sh && ln -s /bin/bash /bin/sh
WORKDIR /output
