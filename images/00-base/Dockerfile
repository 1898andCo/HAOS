### BASE ###
FROM alpine:3.12 as base
ARG ARCH
RUN apk --update --no-cache add \
    bash=5.0.17-r0 \
    bash-completion=2.10-r0 \
    blkid=2.35.2-r0 \
    busybox-initscripts=3.2-r2 \
    ca-certificates=20191127-r4 \
    connman=1.38-r2 \
    conntrack-tools=1.4.6-r0\
    coreutils=8.32-r0 \
    curl=7.77.0-r0 \
    dbus=1.12.18-r0 \
    dmidecode=3.2-r0 \
    dosfstools=4.1-r1 \
    e2fsprogs=1.45.6-r0 \
    e2fsprogs-extra=1.45.6-r0 \
    efibootmgr=17-r0 \
    eudev=3.2.9-r3 \
    findutils=4.7.0-r0 \
    grub-efi=2.04-r1 \
    haveged=1.9.8-r1 \
    htop=2.2.0-r0 \
    hvtools=4.11.9-r1 \
    iproute2=5.6.0-r0 \
    iptables=1.8.4-r2 \
    irqbalance=1.6.0-r1 \
    iscsi-scst=2.2.1-r3 \
    jq=1.6-r1 \
    kbd-bkeymaps=2.2.0-r2\
    lm-sensors=3.6.0-r0 \
    logrotate=3.16.0-r0 \
    lsscsi=0.31-r0 \
    lvm2=2.02.186-r1 \
    lvm2-extra=2.02.186-r1 \
    mdadm=4.1-r2 \
    mdadm-misc=4.1-r2 \
    mdadm-udev=4.1-r2 \
    multipath-tools=0.8.4-r1 \
    ncurses=6.2_p20200523-r0 \
    ncurses-terminfo=6.2_p20200523-r0 \
    nfs-utils=2.4.3-r1 \
    open-iscsi=2.1.2-r0 \
    openrc=0.42.1-r11 \
    openssh-client=8.3_p1-r2 \
    openssh-server=8.3_p1-r2 \
    parted=3.3-r0 \
    procps=3.3.16-r0 \
    qemu-guest-agent=5.0.0-r2 \
    rng-tools=6.10-r2 \
    rsync=3.1.3-r3 \
    strace=5.6-r0 \
    smartmontools=7.1-r2 \
    sudo=1.9.5p2-r0 \
    tar=1.32-r2 \
    tzdata=2021a-r0 \
    util-linux=2.35.2-r0 \
    vim=8.2.0735-r0 \
    wireguard-tools=1.0.20200510-r0 \
    wpa_supplicant=2.9-r8 \
    xfsprogs=5.6.0-r1 \
    xz=5.2.5-r0 \
 && mv -vf /etc/conf.d/qemu-guest-agent /etc/conf.d/qemu-guest-agent.orig \
 && mv -vf /etc/conf.d/rngd             /etc/conf.d/rngd.orig \
 && mv -vf /etc/conf.d/udev-settle      /etc/conf.d/udev-settle.orig \
# replicate the default "no idea, friend" behavior of virt-what
 && touch /usr/sbin/virt-what \
 && chmod +x /usr/sbin/virt-what
