FROM scratch
COPY build/haos/system/ /haos/system/
ENV PATH /haos/system/haos/current:/haos/system/k3s/current:${PATH}
ENTRYPOINT ["haos"]
CMD ["help"]
