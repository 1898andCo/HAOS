ARG REPO
ARG TAG
FROM ${REPO}/k3os-base:${TAG} AS base
RUN apk --update --no-cache add \
    grub-bios=2.04-r1 \
    open-vm-tools=11.1.0-r3 \
    open-vm-tools-deploypkg=11.1.0-r3 \
    open-vm-tools-guestinfo=11.1.0-r3 \
    open-vm-tools-static=11.1.0-r3 \
    open-vm-tools-vmbackup=11.1.0-r3