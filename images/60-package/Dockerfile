ARG REPO
ARG TAG
FROM ${REPO}/blaos-kernel:${TAG} as kernel

ARG REPO
ARG TAG
FROM ${REPO}/blaos-package:${TAG}
ARG VERSION

COPY --from=kernel /output/ /output/BLAOS/system/kernel/

WORKDIR /output/BLAOS/system/kernel
RUN mkdir -vp $(cat version)
RUN ln -sf $(cat version) current
RUN mv -vf initrd kernel.squashfs current/
RUN rm -vf version vmlinuz

# [START] Used in the creation of the virtual image version
RUN mkdir /output-vm/ && \
    mv -vf kernel.vm.squashfs /output-vm/
# [END] Used in the creation of the virtual image version