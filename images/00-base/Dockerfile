### BASE ###
FROM alpine:3.17 as base
ARG ARCH
RUN apk --update --no-cache add \
    bash=5.2.15-r0 \ 
    bash-completion=2.11-r4 \
    blkid=2.38.1-r1 \
    ca-certificates=20230506-r0 \
    connman=1.41-r1 \
    conntrack-tools=1.4.7-r0 \
    coreutils=9.1-r0 \
    curl=8.4.0-r0 \
    dbus=1.14.8-r0 \
    dmidecode=3.4-r0 \
    dosfstools=4.2-r1 \
    e2fsprogs=1.46.6-r0 \
    e2fsprogs-extra=1.46.6-r0 \
    efibootmgr=18-r0 \
    eudev=3.2.11-r4 \
    findutils=4.9.0-r3 \
    grub-efi=2.06-r8 \
    haveged=1.9.18-r0 \
    htop=3.2.1-r1 \
    hvtools=4.11.9-r2 \
    iproute2=6.0.0-r1 \
    iptables=1.8.8-r2 \
    irqbalance=1.9.2-r1 \
    iscsi-scst=2.2.1-r4 \
    jq=1.6-r2 \
    kbd-bkeymaps=2.5.1-r3 \
    lm-sensors=3.6.0-r4 \
    logrotate=3.20.1-r3 \
    lsscsi=0.32-r1 \
    lvm2=2.03.17-r1 \
    lvm2-extra=2.03.17-r1 \
    mdadm=4.2-r1 \
    mdadm-misc=4.2-r1 \
    mdadm-udev=4.2-r1 \
    multipath-tools=0.9.3-r2 \
    ncurses=6.3_p20221119-r1 \
    ncurses-terminfo=6.3_p20221119-r1 \
    nfs-utils=2.6.2-r0 \
    open-iscsi=2.1.8-r0 \
    openrc=0.45.2-r7 \
    openssh-client=9.1_p1-r4 \
    openssh-server=9.1_p1-r4 \
    parted=3.5-r0 \
    procps=3.3.17-r3 \
    qemu-guest-agent=7.1.0-r7 \
    rng-tools=6.15-r1 \
    rsync=3.2.7-r0 \
    smartmontools=7.3-r0 \
    strace=6.0-r0 \
    sudo=1.9.12_p2-r1 \
    tar=1.34-r2 \
    tzdata=2023c-r0 \
    util-linux=2.38.1-r1 \
    vim=9.0.0999-r0 \
    wireguard-tools=1.0.20210914-r1 \
    wpa_supplicant=2.10-r4 \
    xfsprogs=6.0.0-r0 \
    xz=5.2.9-r0

RUN mv -vf /etc/conf.d/qemu-guest-agent /etc/conf.d/qemu-guest-agent.orig && \
    mv -vf /etc/conf.d/rngd /etc/conf.d/rngd.orig && \
    mv -vf /etc/conf.d/udev-settle /etc/conf.d/udev-settle.orig && \
    # replicate the default "no idea, friend" behavior of virt-what
    touch /usr/sbin/virt-what && \
    chmod +x /usr/sbin/virt-what # buildkit
