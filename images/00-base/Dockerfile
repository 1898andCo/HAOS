### BASE ###
FROM alpine:3.16 as base
ARG ARCH
RUN apk --update --no-cache add \
    bash=5.1.16-r2 \ 
    bash-completion=2.11-r4 \
    blkid=2.38-r1 \
    busybox-initscripts=4.2-r0 \
    ca-certificates=20230506-r0 \
    connman=1.41-r1 \
    conntrack-tools=1.4.6-r1 \
    coreutils=9.1-r0\
    curl=8.2.1-r0 \
    dbus=1.14.8-r0 \
    dmidecode=3.3-r1 \
    dosfstools=4.2-r1 \
    e2fsprogs=1.46.6-r0 \
    e2fsprogs-extra=1.46.6-r0 \
    efibootmgr=17-r0 \
    eudev=3.2.11-r0 \
    findutils=4.9.0-r0 \
    grub-efi=2.06-r2 \
    haveged=1.9.18-r0 \
    htop=3.2.0-r1 \
    hvtools=4.11.9-r1 \
    iproute2=5.17.0-r0 \
    iptables=1.8.8-r1 \
    irqbalance=1.8.0-r0 \
    iscsi-scst=2.2.1-r3 \
    jq=1.6-r1 \
    kbd-bkeymaps=2.4.0-r1 \
    lm-sensors=3.6.0-r1 \
    logrotate=3.19.0-r3 \
    lsscsi=0.32-r0 \
    lvm2=2.02.187-r2 \
    lvm2-extra=2.02.187-r2 \
    mdadm=4.2-r0 \
    mdadm-misc=4.2-r0 \
    mdadm-udev=4.2-r0 \
    multipath-tools=0.8.9-r1 \
    ncurses=6.3_p20220521-r1 \
    ncurses-terminfo=6.3_p20220521-r1 \
    nfs-utils=2.6.1-r1 \
    open-iscsi=2.1.6-r0 \
    openrc=0.44.10-r7 \
    openssh-client \
    openssh-server \
    parted=3.5-r0 \
    procps=3.3.17-r1 \
    qemu-guest-agent=7.0.0-r0 \
    rng-tools=6.15-r0 \
    rsync=3.2.7-r0 \
    smartmontools=7.3-r0 \
    strace=5.17-r0 \
    sudo=1.9.12-r1 \
    tar=1.34-r1 \
    tzdata \
    util-linux=2.38-r1 \
    vim=8.2.5000-r0 \
    wireguard-tools=1.0.20210914-r0 \
    wpa_supplicant=2.10-r1 \
    xfsprogs=5.16.0-r1 \
    xz=5.2.5-r1 

RUN mv -vf /etc/conf.d/qemu-guest-agent /etc/conf.d/qemu-guest-agent.orig && \
    mv -vf /etc/conf.d/rngd /etc/conf.d/rngd.orig && \
    mv -vf /etc/conf.d/udev-settle /etc/conf.d/udev-settle.orig && \
    # replicate the default "no idea, friend" behavior of virt-what
    touch /usr/sbin/virt-what && \
    chmod +x /usr/sbin/virt-what # buildkit
