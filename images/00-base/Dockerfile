### BASE ###
FROM alpine:3.18 as base
ARG ARCH
RUN apk --update --no-cache add \
    bash=5.2.15-r5 \ 
    bash-completion=2.11-r5 \
    blkid=2.38.1-r8 \
    ca-certificates=20230506-r0 \
    connman=1.41-r4 \
    conntrack-tools=1.4.7-r1 \
    coreutils=9.3-r1 \
    curl=8.5.0-r0 \
    dbus=1.14.8-r0 \
    dmidecode=3.5-r1 \
    dosfstools=4.2-r2 \
    e2fsprogs=1.47.0-r2 \
    e2fsprogs-extra=1.47.0-r2 \
    efibootmgr=18-r2 \
    eudev=3.2.11-r8 \
    findutils=4.9.0-r5 \
    gcompat=1.1.0-r1 \
    grub-efi \
    haveged=1.9.18-r1 \
    htop=3.2.2-r1 \
    hvtools=4.11.9-r3 \
    iproute2=6.3.0-r0 \
    iptables=1.8.9-r2 \
    irqbalance=1.9.2-r5 \
    iscsi-scst=2.2.1-r5 \
    jq=1.6-r4 \
    kbd-bkeymaps=2.5.1-r7 \
    lm-sensors=3.6.0-r5 \
    logrotate=3.21.0-r1 \
    lsscsi=0.32-r1 \
    lvm2=2.03.21-r3 \
    lvm2-extra=2.03.21-r3 \
    mdadm=4.2-r3 \
    mdadm-misc=4.2-r3 \
    mdadm-udev=4.2-r3 \
    multipath-tools=0.9.5-r1 \
    ncurses=6.4_p20230506-r0 \
    ncurses-terminfo=6.4_p20230506-r0 \
    nfs-utils=2.6.3-r1 \
    open-iscsi=2.1.8-r2 \
    openrc=0.48-r0 \
    openssh-client=9.3_p2-r0 \
    openssh-server=9.3_p2-r0 \
    parted=3.6-r1 \
    procps-ng=4.0.4-r0 \
    qemu-guest-agent=8.0.5-r0\
    rng-tools=6.16-r2 \
    rsync=3.2.7-r4 \
    smartmontools=7.3-r1 \
    strace=6.3-r1 \
    sudo=1.9.13_p3-r2 \
    tar=1.34-r3 \
    tzdata=2023c-r1 \
    util-linux=2.38.1-r8 \
    vim=9.0.2073-r0 \
    virt-what=1.25-r1 \
    wireguard-tools=1.0.20210914-r3 \
    wpa_supplicant=2.10-r7 \
    xfsprogs=6.2.0-r2 \
    xz=5.4.3-r0 \
    grub-bios \
    open-vm-tools=12.3.0-r0 \
    open-vm-tools-deploypkg=12.3.0-r0 \
    open-vm-tools-guestinfo=12.3.0-r0 \
    open-vm-tools-static=12.3.0-r0 \
    open-vm-tools-vmbackup=12.3.0-r0

RUN mv -vf /etc/conf.d/qemu-guest-agent /etc/conf.d/qemu-guest-agent.orig && \
    mv -vf /etc/conf.d/rngd /etc/conf.d/rngd.orig && \
    mv -vf /etc/conf.d/udev-settle /etc/conf.d/udev-settle.orig && \
    # replicate the default "no idea, friend" behavior of virt-what
    touch /usr/sbin/virt-what && \
    chmod +x /usr/sbin/virt-what # buildkit
