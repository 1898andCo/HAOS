ARG REPO
ARG TAG
FROM ${REPO}/blaos-tar:${TAG} as tar

ARG REPO
ARG TAG
FROM ${REPO}/blaos-base:${TAG}
ARG ARCH

COPY --from=tar /output/userspace.tar /output/BLAOS-rootfs-${ARCH}.tar
RUN gzip /output/BLAOS-rootfs-${ARCH}.tar
