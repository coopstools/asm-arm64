package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTokenize(t *testing.T) {
	testCode := ">,28[<+>->+<]."
	cmds := Tokenize(testCode)
	for i, extepctedCmd := range []Cmd{
		{op: INC_IND},
		{op: RD_IN, value: 28},
		{op: CTRL_JMP, value: 8},
		{op: DEC_IND},
		{op: INC_VAL},
		{op: INC_IND},
		{op: DEC_VAL},
		{op: INC_IND},
		{op: INC_VAL},
		{op: DEC_IND},
		{op: CTRL_RTN, value: 8},
	} {
		assert.Equal(t, extepctedCmd, cmds[i])
	}
}

func TestTokenize_subjumps(t *testing.T) {
	testCode := ">[>[-<]<]>[-]+"
	cmds := Tokenize(testCode)
	expectedCmds := []Cmd{
		{op: INC_IND},
		{op: CTRL_JMP, value: 7},
		{op: INC_IND},
		{op: CTRL_JMP, value: 3},
		{op: DEC_VAL}, {op: DEC_IND},
		{op: CTRL_RTN, value: 3},
		{op: DEC_IND},
		{op: CTRL_RTN, value: 7},
		{op: INC_IND},
		{op: CTRL_JMP, value: 2},
		{op: DEC_VAL},
		{op: CTRL_RTN, value: 2},
		{op: INC_VAL},
	}
	require.Equal(t, len(cmds), len(expectedCmds))
	for i, extepctedCmd := range expectedCmds {
		assert.Equal(t, extepctedCmd, cmds[i], "mismatch on line %d", i)
	}
}
