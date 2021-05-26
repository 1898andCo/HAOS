FROM scratch
COPY build/BLAOS/system/ /BLAOS/system/
ENV PATH /BLAOS/system/BLAOS/current:/BLAOS/system/k3s/current:${PATH}
ENTRYPOINT ["BLAOS"]
CMD ["help"]
