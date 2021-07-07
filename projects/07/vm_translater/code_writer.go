package main

import (
	"fmt"
	"io"
	"strconv"
)

type CodeWriter struct {
	w          io.Writer
	fileName   string
	compareNum int
}

// NewCodeWriter returns new CodeWriter.
func NewCodeWriter(w io.Writer) *CodeWriter {
	return &CodeWriter{w, "", 0}
}

// SetFileName sets file name.
func (c *CodeWriter) SetFileName(fileName string) {
	c.fileName = fileName
}

// WriteArithmetic writes Arithmetic.
func (c *CodeWriter) WriteArithmetic(arithmetic string) {
	switch arithmetic {
	case "add":
		// register from top stack
		c.puts("@SP")
		c.puts("M=M-1")
		c.puts("A=M")
		c.puts("D=M")
		c.puts("M=0") // delete top stack
		// add register(=top stack) to second stack
		c.puts("@SP")
		c.puts("M=M-1")
		c.puts("A=M")
		c.puts("M=M+D")
		// increment SP
		c.puts("@SP")
		c.puts("M=M+1")
	case "sub":
		// register from top stack
		c.puts("@SP")
		c.puts("M=M-1")
		c.puts("A=M")
		c.puts("D=M")
		c.puts("M=0") // delete top stack
		// subtract register(=top stack) from second stack
		c.puts("@SP")
		c.puts("M=M-1")
		c.puts("A=M")
		c.puts("M=M-D")
		// increment SP
		c.puts("@SP")
		c.puts("M=M+1")
	case "and":
		// register from top stack
		c.puts("@SP")
		c.puts("M=M-1")
		c.puts("A=M")
		c.puts("D=M")
		c.puts("M=0") // delete top stack
		// second stack & register(=top stack)
		c.puts("@SP")
		c.puts("M=M-1")
		c.puts("A=M")
		c.puts("M=M&D")
		// increment SP
		c.puts("@SP")
		c.puts("M=M+1")
	case "or":
		// register from top stack
		c.puts("@SP")
		c.puts("M=M-1")
		c.puts("A=M")
		c.puts("D=M")
		c.puts("M=0") // delete top stack
		// second stack | register(=top stack)
		c.puts("@SP")
		c.puts("M=M-1")
		c.puts("A=M")
		c.puts("M=M|D")
		// increment SP
		c.puts("@SP")
		c.puts("M=M+1")
	case "neg":
		// negate top stack
		c.puts("@SP")
		c.puts("A=M-1")
		c.puts("M=-M")
	case "not":
		// not top stack
		c.puts("@SP")
		c.puts("A=M-1")
		c.puts("M=!M")
	case "eq", "gt", "lt":
		var m string
		switch arithmetic {
		case "eq":
			m = "JEQ"
		case "gt":
			m = "JGT"
		case "lt":
			m = "JLT"
		}

		// register from top stack
		c.puts("@SP")
		c.puts("M=M-1")
		c.puts("A=M")
		c.puts("D=M")
		c.puts("M=0") // delete top stack
		// subtract register(=top stack) from second stack to register
		c.puts("@SP")
		c.puts("M=M-1")
		c.puts("A=M")
		c.puts("D=M-D")

		// compare
		c.puts("@RETURN_TRUE_" + strconv.Itoa(c.compareNum))
		c.puts("D;" + m)
		// set false
		c.puts("@SP")
		c.puts("A=M")
		c.puts("M=0")
		c.puts("@RETURN_FALSE_" + strconv.Itoa(c.compareNum))
		c.puts("0;JMP")

		// skip setting false when return true
		c.puts("(RETURN_TRUE_" + strconv.Itoa(c.compareNum) + ")")
		// set true
		c.puts("@SP")
		c.puts("A=M")
		c.puts("M=-1")

		// skip setting true when return false
		c.puts("(RETURN_FALSE_" + strconv.Itoa(c.compareNum) + ")")

		// increment SP
		c.puts("@SP")
		c.puts("M=M+1")

		// increment compareNum
		c.compareNum++
	}
}

// WritePushPop writes Push & Pop.
func (c *CodeWriter) WritePushPop(command Command, segment string, index string) {
	if command == C_PUSH {

		switch segment {
		case "constant":
			// write index to register
			c.puts("@" + index)
			c.puts("D=A")
			// push from register
			c.puts("@SP")
			c.puts("A=M")
			c.puts("M=D")
			// increment SP
			c.puts("@SP")
			c.puts("M=M+1")
		}
	}
}

func (c *CodeWriter) puts(mnemonic string) {
	fmt.Fprintln(c.w, mnemonic)
}
