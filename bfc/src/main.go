package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("missing file for compilation")
		os.Exit(1)
	}

	fileName := args[1]
	if len(fileName) < 3 || ".bf" != fileName[len(fileName)-3:] {
		fmt.Println("must supply valid bf file: ", fileName)
		os.Exit(2)
	}

	if len(args) < 3 {
		fmt.Println("missing output file name")
		os.Exit(3)
	}
	outputFileName := args[2]
	cmds := CreateTokensFromFileName(fileName)
	output := InjectTokensAsCode(cmds)

	err := os.WriteFile(outputFileName, []byte(output), 0644)
	if err != nil {
		fmt.Println("could not output c file")
		os.Exit(5)
	}
}
