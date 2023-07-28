ARG REPO
ARG TAG
FROM ${REPO}/haos-base:${TAG} AS base
RUN apk --update --no-cache add \
  grub-bios=2.04-r2 \
  open-vm-tools=11.2.0-r0 \
  open-vm-tools-deploypkg=11.2.0-r0 \
  open-vm-tools-guestinfo=11.2.0-r0 \
  open-vm-tools-static=11.2.0-r0 \
  open-vm-tools-vmbackup=11.2.0-r0