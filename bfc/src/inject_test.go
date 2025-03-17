package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestInject(t *testing.T) {
	cmds := []Cmd{
		{op: INC_IND}, {op: DEC_IND}, {op: INC_VAL}, {op: DEC_VAL},
	}

	output := InjectTokensAsCode(cmds)
	start := strings.Index(output, "// start custom code")
	stop := strings.Index(output, "// stop custom code")
	require.NotEqual(t, -1, start, "missing beginning string")
	require.NotEqual(t, -1, stop, "missing ending string")
	assert.Equal(t, body1, output[start+21:stop])
}

func TestInject_flowControl_multiJmp(t *testing.T) {
	// ><[+<]>[+<]-
	cmds := []Cmd{
		{op: INC_IND}, {op: DEC_IND},
		{op: CTRL_JMP, value: 3}, // JMP 1
		{op: INC_VAL}, {op: DEC_IND},
		{op: CTRL_RTN, value: 3},
		{op: INC_IND},
		{op: CTRL_JMP, value: 3}, // JMP 2
		{op: INC_VAL}, {op: DEC_IND},
		{op: CTRL_RTN, value: 3},
		{op: DEC_VAL},
	}

	output := InjectTokensAsCode(cmds)
	start := strings.Index(output, "// start custom code")
	stop := strings.Index(output, "// stop custom code")
	require.NotEqual(t, -1, start, "missing beginning string")
	require.NotEqual(t, -1, stop, "missing ending string")
	assert.Equal(t, body2, output[start+21:stop])
}

func TestInject_flowControl_subJump(t *testing.T) {
	// >[>[-<]<]>[-]+
	cmds := []Cmd{
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

	output := InjectTokensAsCode(cmds)
	start := strings.Index(output, "// start custom code")
	stop := strings.Index(output, "// stop custom code")
	require.NotEqual(t, -1, start, "missing beginning string")
	require.NotEqual(t, -1, stop, "missing ending string")
	assert.Equal(t, body3, output[start+21:stop])
}

var body1 = `
void f0() {
  movr();
  movl();
  inc();
  dec();
}
`

// ><[+<]>[+<]-
var body2 = `
void f1() {
  inc();
  movl();
}
void f2() {
  inc();
  movl();
}
void f0() {
  movr();
  movl();
  while(ptr[i]!=0){f1();}
  movr();
  while(ptr[i]!=0){f2();}
  dec();
}
`

// >[>[-<]<]>[-]+
var body3 = `
void f11() {
  dec();
  movl();
}
void f1() {
  movr();
  while(ptr[i]!=0){f11();}
  movl();
}
void f2() {
  dec();
}
void f0() {
  movr();
  while(ptr[i]!=0){f1();}
  movr();
  while(ptr[i]!=0){f2();}
  inc();
}
`
