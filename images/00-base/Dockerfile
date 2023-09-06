### BASE ###
FROM alpine:3.16 as base
ARG ARCH
RUN apk --update --no-cache add \
    bash=5.1.16-r0 \ 
    bash-completion=2.11-r2 \
    blkid=2.37.4-r0 \
    busybox-initscripts=3.2-r2 \
    ca-certificates=20220614-r0 \
    connman=1.39-r0 \
    conntrack-tools=1.4.6-r0 \
    coreutils=8.32-r2 \
    curl=7.79.1-r3 \
    dbus=1.12.24-r0 \
    dmidecode=3.3-r0 \
    dosfstools=4.1-r1 \
    e2fsprogs=1.45.7-r0 \
    e2fsprogs-extra=1.45.7-r0 \
    efibootmgr=17-r0 \
    eudev=3.2.9-r3 \
    findutils=4.8.0-r0 \
    grub-efi=2.04-r2 \
    haveged=1.9.14-r1 \
    htop=3.0.5-r1 \
    hvtools=4.11.9-r1 \
    iproute2=5.10.0-r1 \
    iptables=1.8.6-r0 \
    irqbalance=1.7.0-r0 \
    iscsi-scst=2.2.1-r3 \
    jq=1.6-r1 \
    kbd-bkeymaps=2.4.0-r0 \
    lm-sensors=3.6.0-r1 \
    logrotate=3.18.0-r3 \
    lsscsi=0.31-r0 \
    lvm2=2.02.187-r1 \
    lvm2-extra=2.02.187-r1 \
    mdadm=4.1-r3 \
    mdadm-misc=4.1-r3 \
    mdadm-udev=4.1-r3 \
    multipath-tools=0.8.5-r0 \
    ncurses=6.2_p20210109-r1 \
    ncurses-terminfo=6.2_p20210109-r1 \
    nfs-utils=2.5.2-r0 \
    open-iscsi=2.1.3-r1 \
    openrc=0.42.1-r22 \
    openssh-client=8.4_p1-r4 \
    openssh-server=8.4_p1-r4 \
    parted=3.3-r1 \
    procps=3.3.16-r0 \
    qemu-guest-agent=5.2.0-r3 \
    rng-tools=6.10-r2 \
    rsync=3.2.5-r0 \
    strace=5.10-r0 \
    smartmontools=7.2-r1 \
    sudo=1.9.12-r1 \
    tar=1.34-r0 \
    tzdata \
    util-linux=2.37.4-r0 \
    vim=8.2.4836-r0 \
    wireguard-tools=1.0.20200827-r1 \
    wpa_supplicant=2.9-r11 \
    xfsprogs=5.10.0-r0 \
    xz=5.2.5-r1 

RUN mv -vf /etc/conf.d/qemu-guest-agent /etc/conf.d/qemu-guest-agent.orig && \
    mv -vf /etc/conf.d/rngd /etc/conf.d/rngd.orig && \
    mv -vf /etc/conf.d/udev-settle /etc/conf.d/udev-settle.orig && \
    # replicate the default "no idea, friend" behavior of virt-what
    touch /usr/sbin/virt-what && \
    chmod +x /usr/sbin/virt-what # buildkit
