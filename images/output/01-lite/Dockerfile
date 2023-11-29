ARG REPO
ARG TAG
FROM ${REPO}/haos-tar:${TAG} as tar

ARG REPO
ARG TAG
FROM ${REPO}/haos-base:${TAG}
ARG ARCH

COPY --from=tar /output/userspace.tar /output/haos-rootfs-${ARCH}.tar
RUN gzip /output/haos-rootfs-${ARCH}.tar
