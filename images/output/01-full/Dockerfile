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

COPY --from=kernel /output/vmlinuz /output/k3os-vmlinuz-${ARCH}
COPY --from=kernel /output/initrd /output/k3os-initrd-${ARCH}
COPY --from=kernel /output/kernel.squashfs /output/k3os-kernel-${ARCH}.squashfs
COPY --from=kernel /output/version /output/k3os-kernel-version-${ARCH}
COPY --from=iso /output/haos.iso /output/k3os-${ARCH}.iso
COPY --from=tar /output/userspace.tar /output/k3os-rootfs-${ARCH}.tar
RUN gzip /output/k3os-rootfs-${ARCH}.tar
