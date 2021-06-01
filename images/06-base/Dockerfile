ARG REPO
ARG TAG
FROM ${REPO}/k3os-base:${TAG}

ARG ARCH
ENV ARCH ${ARCH}
ENV VERSION v1.43.0
ADD https://github.com/replicatedhq/kots/releases/download/${VERSION}/kots_linux_amd64.tar.gz /usr/local/bin