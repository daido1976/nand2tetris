package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	vmFilePath := os.Args[len(os.Args)-1]
	rf, err := os.Open(vmFilePath)
	if err != nil {
		fmt.Println("Error! The file can not open")
	}

	// initialize parser
	parser := NewParser(rf)

	// create file for writing
	asmFilePath := strings.ReplaceAll(vmFilePath, ".vm", ".asm")
	wf, err := os.Create(asmFilePath)
	if err != nil {
		fmt.Println("Error! The file can not create")
	}
	// initialize code writer
	codeWriter := NewCodeWriter(wf)

	for parser.HasMoreCommands() {
		parser.Advance()

		switch parser.CommandType() {
		case C_ARITHMETIC:
			codeWriter.WriteArithmetic(parser.Arg1())
		case C_PUSH:
			codeWriter.WritePushPop(parser.CommandType(), parser.Arg1(), parser.Arg2())
		}
	}
}
