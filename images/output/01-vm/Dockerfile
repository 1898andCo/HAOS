ARG REPO
ARG TAG
FROM ${REPO}/blaos-iso:${TAG} as iso

ARG REPO
ARG TAG
FROM ${REPO}/blaos-base:${TAG}
ARG ARCH

COPY --from=iso /output/BLAOS-vm.iso /output/BLAOS-vm-${ARCH}.iso