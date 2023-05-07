#!/bin/sh
set -eux

riscv64-linux-gnu-gcc -static -O3 -o hello-riscv64 hello.c

# # To build hello-riscv32:
# git clone https://github.com/riscv/riscv-gnu-toolchain.git
# (
#   cd riscv-gnu-toolchain
#   ./configure --prefix=/opt/riscv --with-arch=rv32ia --with-abi=ilp32
#   make linux
# )

# /opt/riscv/bin/riscv32-unknown-linux-gnu-gcc -march=rv32i -mabi=ilp32 -O3 -static -o hello-riscv32 hello.c
# /opt/riscv/bin/riscv32-unknown-linux-gnu-strip hello-riscv32
