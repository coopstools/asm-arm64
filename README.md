# asm-arm64
Testing writing assembly code in ARM64 in order to create a full compiler

This repo is still developing. Currently, bf code can be compiled into an executable. However, as the structure of this repo is in flux, I will postpone the instructions.

The short of it, for those that don't mind the mess, is that you can run `make` and it will take the file in bf_code and compile it into a binary called "out". You can change the targets using env variables:

```bash
make in=bf_code/t1.bf out=t1
```


