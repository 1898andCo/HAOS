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

COPY --from=kernel /output/kernel.vm.squashfs /output/haos-kernel-vm-${ARCH}.squashfs
COPY --from=iso /output/haos-vm.iso /output/haos-vm-${ARCH}.iso