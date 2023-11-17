ARG REPO
ARG TAG
FROM ${REPO}/k3os-base:${TAG} AS base
RUN apk --update --no-cache add \
    grub-bios=2.06-r8 \
    open-vm-tools=12.1.0-r0 \
    open-vm-tools-deploypkg=12.1.0-r0 \
    open-vm-tools-guestinfo=12.1.0-r0 \
    open-vm-tools-static=12.1.0-r0 \
    open-vm-tools-vmbackup=12.1.0-r0
