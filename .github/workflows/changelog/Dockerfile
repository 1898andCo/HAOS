FROM quay.io/git-chglog/git-chglog:latest as base
RUN git fetch origin --prune --tags
RUN git-chglog ${INPUT_TAG} > RELEASELOG.md