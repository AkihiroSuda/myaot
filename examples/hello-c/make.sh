#!/bin/sh
set -eux

# git clone https://github.com/riscv/riscv-gnu-toolchain.git
# cd riscv-gnu-toolchain
# ./configure --prefix=/opt/riscv --with-arch=rv32ia --with-abi=ilp32
# make linux

/opt/riscv/bin/riscv32-unknown-linux-gnu-gcc -march=rv32i -mabi=ilp32 -O3 -static -o hello-riscv32 hello.c
# /opt/riscv/bin/riscv32-unknown-linux-gnu-strip hello-riscv32
