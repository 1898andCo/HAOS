#!/bin/bash

exec /usr/bin/qemu-system \
    -enable-kvm \
    -nographic \
    -serial mon:stdio \
    -display none \
    -rtc \
    base=utc,clock=host \
    -cdrom /output/k3os.iso \
    -m 2048 \
    -smp 2 \
    -device virtio-rng-pci \
    -net nic \
    -net user,hostfwd=::2222-:22 \
    -drive if=virtio,file=/hd.img

# [START] Used in the creation of the virtual image version
exec /usr/bin/qemu-system \
    -enable-kvm \
    -nographic \
    -serial mon:stdio \
    -display none \
    -rtc \
    base=utc,clock=host \
    -cdrom /output/k3os-vm.iso \
    -m 2048 \
    -smp 2 \
    -device virtio-rng-pci \
    -net nic \
    -net user,hostfwd=::2222-:22 \
    -drive if=virtio,file=/hd-vm.img
# [END] Used in the creation of the virtual image version