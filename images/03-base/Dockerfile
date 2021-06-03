ARG REPO
ARG TAG
FROM ${REPO}/k3os-base:${TAG} AS base
RUN apk --update --no-cache add \
    virt-what=1.20-r0
