ARG REPO
ARG TAG
ARG VERSION
FROM ${REPO}/haos-gobuild:${TAG} as gobuild

ENV LINUXKIT v0.8

FROM gobuild as linuxkit
ENV GO111MODULE off
RUN git clone https://github.com/BOHICA-LABS/linuxkit.git $GOPATH/src/github.com/linuxkit/linuxkit
WORKDIR $GOPATH/src/github.com/linuxkit/linuxkit/pkg/metadata
RUN git checkout develop
RUN gobuild -o /output/metadata

FROM gobuild as haos
ARG VERSION
COPY go.mod $GOPATH/src/github.com/1898andCo/HAOS/
COPY go.sum $GOPATH/src/github.com/1898andCo/HAOS/
COPY /pkg/ $GOPATH/src/github.com/1898andCo/HAOS/pkg/
COPY /main.go $GOPATH/src/github.com/1898andCo/HAOS/
COPY /vendor/ $GOPATH/src/github.com/1898andCo/HAOS/vendor/
WORKDIR $GOPATH/src/github.com/1898andCo/HAOS
RUN gobuild -mod=readonly -o /output/HAOS

FROM gobuild
COPY --from=linuxkit /output/ /output/
COPY --from=haos /output/ /output/
WORKDIR /output
RUN git clone --branch v0.7.0 https://github.com/ahmetb/kubectx.git \
    && chmod -v +x kubectx/kubectx kubectx/kubens
