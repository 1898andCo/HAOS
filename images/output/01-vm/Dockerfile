ARG REPO
ARG TAG
FROM ${REPO}/k3os-iso:${TAG} as iso

ARG REPO
ARG TAG
FROM ${REPO}/k3os-kernel:${TAG} as kernel

ARG REPO
ARG TAG
FROM ${REPO}/haos-base:${TAG}
ARG ARCH

COPY --from=kernel /output/kernel.vm.squashfs /output/k3os-kernel-vm-${ARCH}.squashfs
COPY --from=iso /output/k3os-vm.iso /output/k3os-vm-${ARCH}.iso