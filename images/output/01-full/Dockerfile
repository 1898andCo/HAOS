ARG REPO
ARG TAG
FROM ${REPO}/haos-tar:${TAG} as tar

ARG REPO
ARG TAG
FROM ${REPO}/haos-iso:${TAG} as iso

ARG REPO
ARG TAG
FROM ${REPO}/haos-kernel:${TAG} as kernel

ARG REPO
ARG TAG
FROM ${REPO}/haos-base:${TAG}
ARG ARCH

COPY --from=kernel /output/vmlinuz /output/haos-vmlinuz-${ARCH}
COPY --from=kernel /output/initrd /output/haos-initrd-${ARCH}
COPY --from=kernel /output/kernel.squashfs /output/haos-kernel-${ARCH}.squashfs
COPY --from=kernel /output/version /output/haos-kernel-version-${ARCH}
COPY --from=iso /output/haos.iso /output/haos-${ARCH}.iso
COPY --from=tar /output/userspace.tar /output/haos-rootfs-${ARCH}.tar
RUN gzip /output/haos-rootfs-${ARCH}.tar
