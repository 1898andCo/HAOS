FROM golang:1.17.0-alpine
RUN apk --update add \
    git \
    gcc \
    linux-headers \
    musl-dev \
    make \
    libseccomp \
    libseccomp-dev \
    bash=5.1.16-r0 
COPY gobuild /usr/bin/
RUN rm -f /bin/sh && ln -s /bin/bash /bin/sh
WORKDIR /output
