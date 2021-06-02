ARG REPO
ARG TAG
FROM ${REPO}/k3os-base:${TAG} AS base

ARG ARCH
ENV ARCH ${ARCH}
ENV VERSION v1.43.0
RUN curl -L https://github.com/replicatedhq/kots/releases/download/${VERSION}/kots_linux_amd64.tar.gz > kots_linux_amd64.tar.gz
RUN tar -zxvf /kots_linux_amd64.tar.gz kots
RUN rm -rf /kots_linux_amd64.tar.gz
RUN mv -vf kots /sbin/kubectl-kots