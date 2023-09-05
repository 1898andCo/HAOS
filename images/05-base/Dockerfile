ARG REPO
ARG TAG
FROM ${REPO}/haos-base:${TAG} AS base
RUN apk --update --no-cache add \
    grub-bios \
    open-vm-tools \
    open-vm-tools-deploypkg \
    open-vm-tools-guestinfo \
    open-vm-tools-static \
    open-vm-tools-vmbackup