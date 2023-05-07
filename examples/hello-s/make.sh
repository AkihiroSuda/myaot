#!/bin/sh
set -eux
riscv64-linux-gnu-as -o hello-riscv64.o hello-riscv64.s
riscv64-linux-gnu-ld -o hello-riscv64 hello-riscv64.o
rm -f hello-riscv64.o
riscv64-linux-gnu-strip hello-riscv64

riscv64-linux-gnu-as -march=rv32i -mabi=ilp32 -o hello-riscv32.o hello-riscv32.s
riscv64-linux-gnu-ld -melf32lriscv -o hello-riscv32 hello-riscv32.o
rm -f hello-riscv32.o
riscv64-linux-gnu-strip hello-riscv32
