ARG REPO
ARG TAG
FROM ${REPO}/haos-package:${TAG} as package

ARG REPO
ARG TAG
FROM ${REPO}/haos-base:${TAG} as base
ARG VERSION
ARG ARCH


RUN apk --update add \
    xorriso=1.5.4-r3 \
    grub \
    grub-efi \
    mtools=4.0.43-r1 \
    libvirt \
    qemu-img
RUN if [ "$ARCH" == "amd64" ]; then \
    apk --update add \ 
    qemu-system-x86_64=8.0.5-r0 \
    grub-bios \
    ovmf \
    ;elif [ "$ARCH" == "arm64" ]; then  \
    apk add qemu-system-aarch647.1.0-r7 \
    ;fi
RUN ln -s /usr/bin/qemu-system-* /usr/bin/qemu-system
RUN qemu-img create -f qcow2 /hd.img 40G
COPY run-kvm.sh /usr/bin/
COPY grub.cfg /usr/src/iso/boot/grub/grub.cfg

COPY --from=package /output/ /usr/src/iso/

COPY config.yaml /usr/src/iso/haos/system/
RUN mkdir -p /output && \
    grub-mkrescue -o /output/haos.iso /usr/src/iso/. -- -volid HAOS -joliet on && \
    [ -e /output/haos.iso ] # grub-mkrescue doesn't exit non-zero on failure

# [START] Used in the creation of the virtual image version
RUN rm -f /usr/src/iso/haos/system/kernel/current/kernel.squashfs
COPY --from=package /output-vm/kernel.vm.squashfs /usr/src/iso/haos/system/kernel/current/kernel.squashfs
RUN mkdir -p /output && \
    grub-mkrescue -o /output/haos-vm.iso /usr/src/iso/. -- -volid HAOS -joliet on && \
    [ -e /output/haos-vm.iso ] # grub-mkrescue doesn't exit non-zero on failure
# [END] Used in the creation of the virtual image version


CMD ["run-kvm.sh"]
