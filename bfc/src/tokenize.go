package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type Op int

const (
	INC_IND  Op = 0
	DEC_IND  Op = 1
	INC_VAL  Op = 2
	DEC_VAL  Op = 3
	RD_IN    Op = 4
	RW_OUT   Op = 5
	CTRL_JMP Op = 6
	CTRL_RTN Op = 7

	RW_DEBUG = 8
)

type Cmd struct {
	op    Op
	value int
}

func Tokenize(rawCmd string) []Cmd {
	var cmds []Cmd
	var pStack []int
	var scan []rune
	var p int
	for _, r := range rawCmd {
		if scan != nil {
			if unicode.IsDigit(r) {
				scan = append(scan, r)
				continue
			}
			val := -1
			if len(scan) >= 1 {
				val, _ = strconv.Atoi(string(scan))
			}
			cmds = append(cmds, Cmd{op: RD_IN, value: val % 256})
			scan = nil
			p++
		}
		switch r {
		case '>':
			cmds = append(cmds, Cmd{op: INC_IND})
			p++
		case '<':
			cmds = append(cmds, Cmd{op: DEC_IND})
			p++
		case '+':
			cmds = append(cmds, Cmd{op: INC_VAL})
			p++
		case '-':
			cmds = append(cmds, Cmd{op: DEC_VAL})
			p++
		case ',':
			scan = []rune{}
		case '.':
			cmds = append(cmds, Cmd{op: RW_OUT})
			p++
		case '[':
			cmds = append(cmds, Cmd{op: CTRL_JMP})
			pStack = append(pStack, p)
			p++
		case ']':
			lastP := pStack[len(pStack)-1]
			pStack = pStack[:len(pStack)-1]
			cmds = append(cmds, Cmd{op: CTRL_RTN, value: p - lastP})
			cmds[lastP].value = p - lastP
			p++
		case '#':
			if len(cmds) > 0 && cmds[len(cmds)-1].op == RW_DEBUG {
				cmds[len(cmds)-1].value += 1
				break
			}
			cmds = append(cmds, Cmd{op: RW_DEBUG, value: 1})
		}
	}
	// TODO: Add check for mismatch parens
	// TODO: Add proper error handling to avoid breaking REPL
	// TODO: Use compiler in core code
	// TODO: Add stack dumps to compiler
	// TODO: Add stack dumps to Readme
	return cmds
}

func CreateTokensFromFileName(fileName string) []Cmd {
	contents, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("could not open file: ", fileName, err.Error())
		os.Exit(4)
	}
	return Tokenize(string(contents))
}
