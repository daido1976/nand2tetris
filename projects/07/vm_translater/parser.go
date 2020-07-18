package main

import (
	"bufio"
	"io"
	"strings"
)

// Parser parses assembly.
type Parser struct {
	scanner        *bufio.Scanner
	currentCommand string
}

// Command represents command type.
type Command int

const (
	C_NONE Command = iota
	C_ARITHMETIC
	C_PUSH
	C_POP
	C_LABEL
	C_GOTO
	C_IF
	C_FUNCTION
	C_RETURN
	C_CALL
)

// NewParser returns new Parser.
func NewParser(r io.Reader) *Parser {
	s := bufio.NewScanner(r)
	return &Parser{s, ""}
}

// HasMoreCommands load one line and returns whether the command is still present in the input.
func (p *Parser) HasMoreCommands() bool {
	return p.scanner.Scan()
}

// Advance loads the next command and make it the current command.
// It also removes comments and spaces.
func (p *Parser) Advance() {
	p.currentCommand = p.scanner.Text()
	if len(p.currentCommand) <= 0 {
		return
	}

	// remove comments
	tokens := strings.SplitN(p.currentCommand, "//", 2)
	if len(tokens) > 0 {
		p.currentCommand = tokens[0]
	}

	// remove spaces
	p.currentCommand = strings.TrimSpace(p.currentCommand)
}

// CommandType returns current command type.
func (p *Parser) CommandType() Command {
	if p.currentCommand == "" {
		// empty command
		return C_NONE
	}

	switch {
	case strings.HasPrefix(p.currentCommand, "push"):
		return C_PUSH
	case strings.HasPrefix(p.currentCommand, "pop"):
		return C_POP
	case strings.HasPrefix(p.currentCommand, "label"):
		return C_LABEL
	case strings.HasPrefix(p.currentCommand, "goto"):
		return C_GOTO
	case strings.HasPrefix(p.currentCommand, "if-goto"):
		return C_IF
	case strings.HasPrefix(p.currentCommand, "function"):
		return C_FUNCTION
	case strings.HasPrefix(p.currentCommand, "return"):
		return C_RETURN
	case strings.HasPrefix(p.currentCommand, "call"):
		return C_CALL
	}
	return C_ARITHMETIC
}

// Arg1 returns first arg.
func (p *Parser) Arg1() string {
	switch p.CommandType() {
	case C_RETURN:
		panic(`Invalid CommandType`)
	case C_ARITHMETIC:
		return p.currentCommand
	default:
		tokens := strings.Split(p.currentCommand, " ")
		return tokens[1]
	}
}

// Arg2 returns second arg.
func (p *Parser) Arg2() string {
	switch p.CommandType() {
	case C_PUSH, C_POP, C_FUNCTION, C_CALL:
		tokens := strings.Split(p.currentCommand, " ")
		return tokens[2]
	default:
		panic(`Invalid CommandType`)
	}
}
