ARG REPO
ARG TAG
FROM ${REPO}/haos-rootfs:${TAG} as rootfs

ARG REPO
ARG TAG
FROM ${REPO}/haos-progs:${TAG} as progs

ARG REPO
ARG TAG
FROM ${REPO}/haos-base:${TAG}

COPY --from=rootfs /output/rootfs.squashfs /usr/src/
COPY install.sh /output/haos-install.sh
COPY --from=progs /output/HAOS /output/HAOS
RUN echo -n "_sqmagic_" >> /output/HAOS
RUN cat /usr/src/rootfs.squashfs >> /output/HAOS
