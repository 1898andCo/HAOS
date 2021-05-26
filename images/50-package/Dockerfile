ARG REPO
ARG TAG
FROM ${REPO}/blaos-k3s:${TAG} as k3s

ARG REPO
ARG TAG
FROM ${REPO}/blaos-bin:${TAG} as bin

ARG REPO
ARG TAG
FROM ${REPO}/blaos-base:${TAG} as base
ARG VERSION

COPY --from=k3s /output/  /output/BLAOS/system/k3s/
COPY --from=bin /output/  /output/BLAOS/system/BLAOS/${VERSION}/

WORKDIR /output/BLAOS/system/k3s
RUN mkdir -vp $(cat version) /output/sbin
RUN mv -vf crictl ctr kubectl /output/sbin/
RUN ln -sf $(cat version) current
RUN mv -vf install.sh current/k3s-install.sh
RUN mv -vf k3s current/
RUN rm -vf version *.sh
RUN ln -sf /BLAOS/system/k3s/current/k3s /output/sbin/k3s

WORKDIR /output/BLAOS/system/BLAOS
RUN ln -sf ${VERSION} current
RUN ln -sf /BLAOS/system/BLAOS/current/BLAOS /output/sbin/BLAOS
RUN ln -sf BLAOS /output/sbin/init
