package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestInject(t *testing.T) {
	cmds := []Cmd{
		{op: INC_IND}, {op: DEC_IND}, {op: INC_VAL}, {op: DEC_VAL},
	}

	output := Inject(cmds)
	start := strings.Index(output, "int main(")
	assert.NotEqual(t, -1, start)
	assert.Equal(t, body, output[start:])
}

var body = `int main(void) {
  i = 0;
  ptr = (int *)calloc(sizeof(int), 16384);
  movr();
  movl();
  inc();
  dec();

  for (int ii = 0; ii <= 6; ii++)
    printf("%d ", ptr[ii]);
  free(ptr);
  return 0;
}
`
