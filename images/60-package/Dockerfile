ARG REPO
ARG TAG
FROM ${REPO}/haos-kernel:${TAG} as kernel

ARG REPO
ARG TAG
FROM ${REPO}/haos-package:${TAG}
ARG VERSION

COPY --from=kernel /output/ /output/haos/system/kernel/

WORKDIR /output/haos/system/kernel
RUN mkdir -vp $(cat version)
RUN ln -sf $(cat version) current
RUN mv -vf initrd kernel.squashfs current/
RUN rm -vf version vmlinuz

# [START] Used in the creation of the virtual image version
RUN mkdir /output-vm/ && \
    mv -vf kernel.vm.squashfs /output-vm/
# [END] Used in the creation of the virtual image version