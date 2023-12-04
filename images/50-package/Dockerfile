ARG REPO
ARG TAG
FROM ${REPO}/haos-k3s:${TAG} as k3s

ARG REPO
ARG TAG
FROM ${REPO}/haos-bin:${TAG} as bin

ARG REPO
ARG TAG
FROM ${REPO}/haos-base:${TAG} as base
ARG VERSION

COPY --from=k3s /output/  /output/haos/system/k3s/
COPY --from=bin /output/  /output/haos/system/haos/${VERSION}/

WORKDIR /output/haos/system/k3s
RUN mkdir -vp $(cat version) /output/sbin
RUN mv -vf crictl ctr kubectl /output/sbin/
RUN ln -sf $(cat version) current
RUN mv -vf install.sh current/k3s-install.sh
RUN mv -vf k3s current/
RUN rm -vf version *.sh
RUN ln -sf /haos/system/k3s/current/k3s /output/sbin/k3s

WORKDIR /output/haos/system/haos
RUN ln -sf ${VERSION} current
RUN ln -sf /haos/system/haos/current/haos /output/sbin/haos
RUN ln -sf haos /output/sbin/init
