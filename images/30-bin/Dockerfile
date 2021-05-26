ARG REPO
ARG TAG
FROM ${REPO}/blaos-rootfs:${TAG} as rootfs

ARG REPO
ARG TAG
FROM ${REPO}/blaos-progs:${TAG} as progs

ARG REPO
ARG TAG
FROM ${REPO}/blaos-base:${TAG}

COPY --from=rootfs /output/rootfs.squashfs /usr/src/
COPY install.sh /output/BLAOS-install.sh
COPY --from=progs /output/BLAOS /output/BLAOS
RUN echo -n "_sqmagic_" >> /output/BLAOS
RUN cat /usr/src/rootfs.squashfs >> /output/BLAOS
