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

func Inject(cmds []Cmd) string {
	code := ""
	for _, cmd := range cmds {
		if v, ok := cmdLookup[cmd.op]; ok {
			code = fmt.Sprintf("%s  %s\n", code, v)
		}
	}
	return strings.Replace(template, "{{$}}", code, 1)
}
