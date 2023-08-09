ARG REPO
ARG TAG
FROM ${REPO}/k3os-base:${TAG} AS base

ARG ARCH
ENV ARCH ${ARCH}
ENV VERSION v1.101.0
ADD  ./kots /sbin/kubectl-kots

