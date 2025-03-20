# asm-arm64
Testing writing assembly code in ARM64 in order to create a full compiler. This repo is for testing and messing around. For an actual compiler, checkout [BrainF**k](https://github.com/coopstools/brainf-k).

This repo is still developing. Currently, bf code can be compiled into an executable. However, as the structure of this repo is in flux, I will postpone the instructions.

The short of it, for those that don't mind the mess, is that you can run `make` and it will take the file in bf_code and compile it into a binary called "out". You can change the targets using env variables:

```bash
$ make in=bf_code/t1.bf out=t1
```

One of the code examples available is a sorting algorithm written in BF. The args for any compiled application can be delimited with any non-numeric character.

```bash
$ make in=bf_code/sort.bf out=sort
$ ./sort 23,45,16,81,71
81 71 45 23 16
$ ./sort 23 45 16 81 71
81 71 45 23 16
$ ./sort 23%e45%e16%e81%e71
81 71 45 23 16 
```

## TODO

- [ ] Make the size of inputs variable
- [ ] Move to standard input when size of inputs exceeded
