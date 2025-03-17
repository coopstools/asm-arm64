package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed template.txt
var template string

var cmdLookup = map[Op]string{
	INC_IND: "movr();",
	DEC_IND: "movl();",
	INC_VAL: "inc();",
	DEC_VAL: "dec();",
}

func inject(cmds []Cmd, depth int) string {
	var subfuncs []string
	code := ""
	subcount := 1
	for i := 0; i < len(cmds); i++ {
		cmd := cmds[i]
		if cmd.op == CTRL_JMP {
			subfunc := inject(cmds[i+1:i+cmd.value], depth*10+subcount)
			subfuncs = append(subfuncs, subfunc)
			i = i + cmd.value
			code = fmt.Sprintf("%s  while(ptr[i]!=0){f%d();}\n", code, depth*10+subcount)
			subcount += 1
			continue
		}
		if v, ok := cmdLookup[cmd.op]; ok {
			code = fmt.Sprintf("%s  %s\n", code, v)
		}
	}
	subfuncsJoined := strings.Join(subfuncs, "")
	return fmt.Sprintf("%s\nvoid f%d() {\n%s}", subfuncsJoined, depth, code)
}

func InjectTokensAsCode(cmds []Cmd) string {
	code := inject(cmds, 0)
	return strings.Replace(template, "{{$funcs}}", code, 1)
}
