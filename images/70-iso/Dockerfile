ARG REPO
ARG TAG
FROM ${REPO}/blaos-package:${TAG} as package

ARG REPO
ARG TAG
FROM ${REPO}/blaos-base:${TAG} as base
ARG VERSION
ARG ARCH
RUN apk add xorriso grub grub-efi mtools libvirt qemu-img
RUN if [ "$ARCH" == "amd64" ]; then \
        apk add qemu-system-x86_64 grub-bios ovmf \
    ;elif [ "$ARCH" == "arm64" ]; then \
        apk add qemu-system-aarch64 \
    ;fi
RUN ln -s /usr/bin/qemu-system-* /usr/bin/qemu-system
RUN qemu-img create -f qcow2 /hd.img 40G
COPY run-kvm.sh /usr/bin/
COPY grub.cfg /usr/src/iso/boot/grub/grub.cfg

COPY --from=package /output/ /usr/src/iso/

COPY config.yaml /usr/src/iso/BLAOS/system/
RUN mkdir -p /output && \
    grub-mkrescue -o /output/BLAOS.iso /usr/src/iso/. -- -volid BLAOS -joliet on && \
    [ -e /output/BLAOS.iso ] # grub-mkrescue doesn't exit non-zero on failure

# [START] Used in the creation of the virtual image version
RUN rm -f /usr/src/iso/BLAOS/system/kernel/current/kernel.squashfs
COPY --from=package /output-vm/kernel.vm.squashfs /usr/src/iso/BLAOS/system/kernel/current/kernel.squashfs
RUN mkdir -p /output && \
    grub-mkrescue -o /output/BLAOS-vm.iso /usr/src/iso/. -- -volid BLAOS -joliet on && \
    [ -e /output/BLAOS-vm.iso ] # grub-mkrescue doesn't exit non-zero on failure
# [END] Used in the creation of the virtual image version


CMD ["run-kvm.sh"]
