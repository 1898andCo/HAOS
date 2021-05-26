ARG REPO
ARG TAG
FROM ${REPO}/blaos-package:${TAG} as package

ARG REPO
ARG TAG
FROM ${REPO}/blaos-base:${TAG} as base
ARG VERSION

COPY --from=package /output/   /usr/src/${VERSION}/
WORKDIR /output
RUN tar cvf userspace.tar -C /usr/src ${VERSION}
