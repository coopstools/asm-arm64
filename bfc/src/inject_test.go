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
	assert.Equal(t, body, output[start+21:stop])
}

var body = `
void f0() {
  movr();
  movl();
  inc();
  dec();
}
`
