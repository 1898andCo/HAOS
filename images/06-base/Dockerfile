ARG REPO
ARG TAG
FROM ${REPO}/haos-base:${TAG} AS base

ARG ARCH
ENV ARCH ${ARCH}
ENV VERSION v1.101.0
ADD  ./kots /sbin/kubectl-kots

