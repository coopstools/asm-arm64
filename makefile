C_FILE ?= csrc/hello.c
FILE ?= src/hello.s
BIN_FILE ?= hello

in ?= bf_code/t1.bf
out ?= out

.PHONY: default
default: precompile compile

init-setup:
	xcode-select --install
	brew install gcc

precompile:
	(cd bfc && go build -o compile ./src)

compile:
	./bfc/compile $(in) $(out).c
	gcc $(out).c -o $(out)
	@rm $(out).c

builda:
	@as $(FILE) -o hello.o
	@ld hello.o -o $(BIN_FILE) -l System -syslibroot `xcrun -sdk macosx --show-sdk-path` -e _main -arch arm64
	@rm hello.o

runa: build
	./$(BIN_FILE) $1
	@rm $(BIN_FILE)

buildc:
	@gcc $(C_FILE) -o $(BIN_FILE)

runc: buildc
	./$(BIN_FILE)
	@rm $(BIN_FILE)



