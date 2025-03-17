C_FILE ?= csrc/hello.c
FILE ?= src/hello.s
BIN_FILE ?= hello

init-setup:
	xcode-select --install
	brew install gcc

build:
	@as $(FILE) -o hello.o
	@ld hello.o -o $(BIN_FILE) -l System -syslibroot `xcrun -sdk macosx --show-sdk-path` -e _main -arch arm64
	@rm hello.o

run: build
	./$(BIN_FILE)
	@rm $(BIN_FILE)

buildc:
	@gcc $(C_FILE) -o $(BIN_FILE)

runc: buildc
	./$(BIN_FILE)
	@rm $(BIN_FILE)

compile:
	./bfc/compile $(in) $(out).c
	gcc $(out).c -o $(out)
	rm $(out).c
