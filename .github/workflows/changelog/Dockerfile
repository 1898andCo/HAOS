FROM quay.io/git-chglog/git-chglog:latest as base
RUN git fetch origin --prune --tags
RUN git-chglog ${DRONE_TAG} > RELEASELOG.md