ARG REPO
ARG TAG
FROM ubuntu:focal

RUN apt-get --assume-yes update \
    && apt-get --assume-yes install \
    curl=7.68.0-1ubuntu2.19 \
    initramfs-tools=0.136ubuntu6.7\
    kmod=27-1ubuntu2.1 \
    lz4=1.9.2-2ubuntu0.20.04.1  \
    rsync=3.1.3-8ubuntu0.5 \
    xz-utils=5.2.4-1ubuntu1.1 \
    && echo 'r8152' >> /etc/initramfs-tools/modules \
    && echo 'hfs' >> /etc/initramfs-tools/modules \
    && echo 'hfsplus' >> /etc/initramfs-tools/modules \
    && echo 'nls_utf8' >> /etc/initramfs-tools/modules \
    && echo 'nls_iso8859_1' >> /etc/initramfs-tools/modules

ARG ARCH
ENV KVERSION=5.4.0-73-generic
ENV URL=https://github.com/1898andCo/HAOS-kernel/releases/tag/0.1

ENV KERNEL_XZ=${URL}/kernel-generic_${ARCH}.tar.xz
ENV KERNEL_EXTRA_XZ=${URL}/kernel-extra-generic_${ARCH}.tar.xz
ENV KERNEL_HEADERS_XZ=${URL}/kernel-headers-generic_${ARCH}.tar.xz

# Download kernel
RUN mkdir -p /usr/src
RUN curl -fL $KERNEL_XZ -o /usr/src/kernel.tar.xz
RUN curl -fL $KERNEL_EXTRA_XZ -o /usr/src/kernel-extra.tar.xz
RUN curl -fL $KERNEL_HEADERS_XZ -o /usr/src/kernel-headers.tar.xz

# Extract to /usr/src/root
RUN mkdir -p /usr/src/root && \
    cd /usr/src/root && \
    tar xvf /usr/src/kernel.tar.xz && \
    tar xvf /usr/src/kernel-extra.tar.xz && \
    tar xvf /usr/src/kernel-headers.tar.xz

# Create initrd
RUN mkdir /usr/src/initrd && \
    rsync -a /usr/src/root/lib/ /lib/ && \
    depmod $KVERSION && \
    mkinitramfs -k $KVERSION -c lz4 -o /usr/src/initrd.tmp

# Generate initrd firmware and module lists
RUN mkdir -p /output/lib && \
    mkdir -p /output/headers && \
    cd /usr/src/initrd && \
    lz4cat /usr/src/initrd.tmp | cpio -idmv && \
    find lib/modules -name \*.ko > /output/initrd-modules && \
    echo lib/modules/${KVERSION}/modules.order >> /output/initrd-modules && \
    echo lib/modules/${KVERSION}/modules.builtin >> /output/initrd-modules && \
    find lib/firmware -type f > /output/initrd-firmware && \
    find usr/lib/firmware -type f | sed 's!usr/!!' >> /output/initrd-firmware

# Copy output assets
RUN cd /usr/src/root && \
    cp -r usr/src/linux-headers* /output/headers && \
    cp -r lib/firmware /output/lib/firmware && \
    cp -r lib/modules /output/lib/modules && \
    cp boot/System.map* /output/System.map && \
    cp boot/config* /output/config && \
    cp boot/vmlinuz-* /output/vmlinuz && \
    echo ${KVERSION} > /output/version
