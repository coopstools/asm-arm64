package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTokenize(t *testing.T) {
	testCode := ">,28[<+>->+<]."
	cmds := Tokenize(testCode)
	for i, extepctedCmd := range []Cmd{
		{op: INC_IND},
		{op: RD_IN, value: 28},
		{op: CTRL_JMP, value: 10},
		{op: DEC_IND},
		{op: INC_VAL},
		{op: INC_IND},
		{op: DEC_VAL},
		{op: INC_IND},
		{op: INC_VAL},
		{op: DEC_IND},
		{op: CTRL_RTN, value: 2},
	} {
		assert.Equal(t, extepctedCmd, cmds[i])
	}
}
