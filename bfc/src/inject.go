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

func inject(cmds []Cmd, depth, start, stop int) string {
	var subfuncs []string
	code := ""
	for _, cmd := range cmds {
		if v, ok := cmdLookup[cmd.op]; ok {
			code = fmt.Sprintf("%s  %s\n", code, v)
		}
	}
	subfuncsJoined := strings.Join(subfuncs, "\n")
	return fmt.Sprintf("%s\nvoid f%d() {\n%s}", subfuncsJoined, depth, code)
}

func InjectTokensAsCode(cmds []Cmd) string {
	code := inject(cmds, 0, 0, len(cmds)-1)
	return strings.Replace(template, "{{$funcs}}", code, 1)
}
