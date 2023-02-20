# MyAOT: An experimental AOT(-ish) compiler (Linux/riscv32 ELF -> Linux/x86\_64 ELF, Darwin/arm64 Mach-O, WASM, ...)

MyAOT is an experimental AOT(-ish) compiler that translates a Linux/riscv32 ELF binary to:
- Linux/x86\_64 ELF
- Darwin/arm64 Mach-O
- WASM
- Or basically whatever, by using C as an intermediate language

## Status
Only "Hello, world" works.

## Usage

For Linux, Darwin, etc:
```console
$ file examples/hello-s/hello-riscv32
examples/hello-s/hello-riscv32: ELF 32-bit LSB executable, UCB RISC-V, soft-float ABI, version 1 (SYSV), statically linked, stripped

$ go install ./cmd/myaot

$ myaot compile ./examples/hello-s/hello-riscv32
INFO[0000] Compiling ./examples/hello-s/hello-riscv32 --> a.out.c
INFO[0000] Compiling a.out.c --> a.out
INFO[0000] Removing a.out.c
INFO[0000] Done: a.out

$ file a.out
a.out: ELF 64-bit LSB pie executable, x86-64, version 1 (SYSV), dynamically linked, interpreter /lib64/ld-linux-x86-64.so.2, BuildID[sha1]=2d7a7211b6b8be795a2a9837bd39a8e1130df642, for GNU/Linux 3.2.0, not stripped

$ ./a.out
Hello World!
```

For WASM:
```console
$ myaot compile -o a.c ./examples/hello-s/hello-riscv32
INFO[0000] Compiling ./examples/hello-s/hello-riscv32 --> a.c
INFO[0000] Done: a.c

$ emcc -o a.wasm a.c

$ file a.wasm
a.wasm: WebAssembly (wasm) binary module version 0x1 (MVP)

$ emcc -o a.html -s WASM=1 a.c

$ emrun a.html
```
