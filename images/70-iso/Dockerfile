ARG REPO
ARG TAG
FROM ${REPO}/k3os-package:${TAG} as package

ARG REPO
ARG TAG
FROM ${REPO}/haos-base:${TAG} as base
ARG VERSION
ARG ARCH
RUN apk --update add \
  xorriso=1.5.2-r0 \
  grub=2.04-r2 \
  grub-efi=2.04-r2 \
  mtools=4.0.26-r0 \
  libvirt=6.10.0-r1 \
  qemu-img=5.2.0-r3
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

COPY config.yaml /usr/src/iso/k3os/system/
RUN mkdir -p /output && \
  grub-mkrescue -o /output/k3os.iso /usr/src/iso/. -- -volid K3OS -joliet on && \
  [ -e /output/k3os.iso ] # grub-mkrescue doesn't exit non-zero on failure

# [START] Used in the creation of the virtual image version
RUN rm -f /usr/src/iso/k3os/system/kernel/current/kernel.squashfs
COPY --from=package /output-vm/kernel.vm.squashfs /usr/src/iso/k3os/system/kernel/current/kernel.squashfs
RUN mkdir -p /output && \
  grub-mkrescue -o /output/k3os-vm.iso /usr/src/iso/. -- -volid K3OS -joliet on && \
  [ -e /output/k3os-vm.iso ] # grub-mkrescue doesn't exit non-zero on failure
# [END] Used in the creation of the virtual image version


CMD ["run-kvm.sh"]
