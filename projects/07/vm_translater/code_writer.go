package main

import (
	"fmt"
	"io"
)

type CodeWriter struct {
	w        io.Writer
	fileName string
}

// NewCodeWriter returns new CodeWriter.
func NewCodeWriter(w io.Writer) *CodeWriter {
	return &CodeWriter{w, ""}
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
		// add register(=top stack) to second stack
		c.puts("@SP")
		c.puts("M=M-1")
		c.puts("A=M")
		c.puts("M=D+M")
		// increment SP
		c.puts("@SP")
		c.puts("M=M+1")
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
