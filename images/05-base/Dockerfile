ARG REPO
ARG TAG
FROM ${REPO}/haos-base:${TAG} AS base
RUN apk --update --no-cache add \
    grub-bios \
    open-vm-tools=12.3.0-r0 \
    open-vm-tools-deploypkg=12.3.0-r0 \
    open-vm-tools-guestinfo=12.3.0-r0 \
    open-vm-tools-static=12.3.0-r0 \
    open-vm-tools-vmbackup=12.3.0-r0
