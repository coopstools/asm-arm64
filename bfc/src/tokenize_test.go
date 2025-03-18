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

func TestTokenize_debug(t *testing.T) {
	testCode := ">#,28>>###"
	cmds := Tokenize(testCode)
	expectedCmds := []Cmd{
		{op: INC_IND},
		{op: RW_DEBUG, value: 1},
		{op: RD_IN, value: 28},
		{op: INC_IND},
		{op: INC_IND},
		{op: RW_DEBUG, value: 3},
	}
	require.Equal(t, len(expectedCmds), len(cmds))
	for i, expectedCmd := range expectedCmds {
		assert.Equal(t, expectedCmd, cmds[i])
	}
}

func TestTokenize_inputs(t *testing.T) {
	testCode := ">,28>,>"
	cmds := Tokenize(testCode)
	expectedCmds := []Cmd{
		{op: INC_IND},
		{op: RD_IN, value: 28},
		{op: INC_IND},
		{op: RD_IN, value: -1},
		{op: INC_IND},
	}
	require.Equal(t, len(expectedCmds), len(cmds))
	for i, expectedCmd := range expectedCmds {
		assert.Equal(t, expectedCmd, cmds[i])
	}
}

// TestTokenize_debugAndFlowControl This test was created to catch a bug with the debugging line messing up the
// stack pointer
func TestTokenize_debugAndFlowControl(t *testing.T) {
	testCode := "some text[>###[,28\ntesting sentence;\n>,>]]"
	cmds := Tokenize(testCode)
	expectedCmds := []Cmd{
		{op: CTRL_JMP, value: 9},
		{op: INC_IND},
		{op: RW_DEBUG, value: 3},
		{op: CTRL_JMP, value: 5},
		{op: RD_IN, value: 28},
		{op: INC_IND},
		{op: RD_IN, value: -1},
		{op: INC_IND},
		{op: CTRL_RTN, value: 5},
		{op: CTRL_RTN, value: 9},
	}
	require.Equal(t, len(expectedCmds), len(cmds))
	for i, expectedCmd := range expectedCmds {
		require.Equal(t, expectedCmd, cmds[i], "error on command %d", i)
	}
}
