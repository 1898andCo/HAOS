ARG REPO
ARG TAG
FROM ${REPO}/blaos-tar:${TAG} as tar

ARG REPO
ARG TAG
FROM ${REPO}/blaos-iso:${TAG} as iso

ARG REPO
ARG TAG
FROM ${REPO}/blaos-kernel:${TAG} as kernel

ARG REPO
ARG TAG
FROM ${REPO}/blaos-base:${TAG}
ARG ARCH

COPY --from=kernel /output/vmlinuz /output/BLAOS-vmlinuz-${ARCH}
COPY --from=kernel /output/initrd /output/BLAOS-initrd-${ARCH}
COPY --from=kernel /output/kernel.squashfs /output/BLAOS-kernel-${ARCH}.squashfs
COPY --from=kernel /output/version /output/BLAOS-kernel-version-${ARCH}
COPY --from=iso /output/BLAOS.iso /output/BLAOS-${ARCH}.iso
COPY --from=tar /output/userspace.tar /output/BLAOS-rootfs-${ARCH}.tar
RUN gzip /output/BLAOS-rootfs-${ARCH}.tar
