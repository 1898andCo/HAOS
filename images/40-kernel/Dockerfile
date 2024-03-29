ARG REPO
ARG TAG
FROM ${REPO}/haos-bin:${TAG} as bin

FROM ${REPO}/haos-kernel-stage1:${TAG} as kernel

FROM ${REPO}/haos-base:${TAG}
ARG TAG
RUN apk --update add \
    squashfs-tools=4.6.1-r1
COPY --from=kernel /output/ /usr/src/kernel/

RUN mkdir -p /usr/src/initrd/lib && \
    cd /usr/src/kernel && \
    tar cf - -T initrd-modules -T initrd-firmware | tar xf - -C /usr/src/initrd/ && \
    depmod -b /usr/src/initrd $(cat /usr/src/kernel/version)

RUN mkdir -p /output && \
    cd /usr/src/kernel && \
    depmod -b . $(cat /usr/src/kernel/version) && \
    mksquashfs . /output/kernel.squashfs

# [START] Used in the creation of the virtual image version
RUN cd /usr/src/kernel && \
    rm -rf headers/* && \
    rm -rf lib/firmware/* && \
    mksquashfs . /output/kernel.vm.squashfs
# [END] Used in the creation of the virtual image version

RUN cp /usr/src/kernel/version /output/ && \
    cp /usr/src/kernel/vmlinuz /output/

COPY --from=bin /output/ /usr/src/haos/
RUN cd /usr/src/initrd && \
    mkdir -p haos/system/haos/${TAG} && \
    cp /usr/src/haos/haos haos/system/haos/${TAG} && \
    ln -s ${TAG} haos/system/haos/current && \
    ln -s /haos/system/haos/current/haos init

RUN cd /usr/src/initrd && \
    find . | cpio -H newc -o | gzip -c -1 > /output/initrd
