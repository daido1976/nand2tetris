package main

import (
	"bufio"
	"strings"
)

type Parser struct {
	scanner        *bufio.Scanner
	currentCommand string
}

type Command int

const (
	N_COMMAND Command = iota // 0
	A_COMMAND                // 1
	C_COMMAND                // 2
	L_COMMAND                // 3
)

// NewParser returns new Parser.
func NewParser(s *bufio.Scanner) *Parser {
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
		return N_COMMAND
	}
	if p.currentCommand[0] == '@' {
		// @Xxx
		return A_COMMAND
	}
	if p.currentCommand[0] == '(' {
		// (Xxx)
		return L_COMMAND
	}
	// dest=comp;jump
	return C_COMMAND
}

// Symbol retunrs @Xxx -> Xxx or (Xxx) -> Xxx.
func (p *Parser) Symbol() string {
	if p.CommandType() == A_COMMAND {
		// A_COMMAND: @Xxx -> Xxx
		return strings.TrimLeft(p.currentCommand, "@")
	}
	// L_COMMAND: (Xxx) -> Xxx
	return strings.TrimRight(strings.TrimLeft(p.currentCommand, "("), ")")
}

// Dest returns dest mnemonic in the C-Instruction.
func (p *Parser) Dest() string {
	tokens := strings.Split(p.currentCommand, "=")
	if len(tokens) <= 1 {
		return ""
	}
	return tokens[0]
}

// Comp returns comp mnemonic in the C-Instruction.
func (p *Parser) Comp() string {
	tokens := strings.Split(p.currentCommand, "=")
	if len(tokens) <= 1 {
		// comp;jump
		t := strings.Split(tokens[0], ";")
		return t[0]
	}
	// dest=comp;jump or dest=comp -> [dest comp;jump] or [dest comp]
	t := strings.Split(tokens[1], ";")
	// t == [comp jump] or [comp]
	return t[0]
}

// Jump returns jump mnemonic in the C-Instruction.
func (p *Parser) Jump() string {
	tokens := strings.Split(p.currentCommand, ";")
	if len(tokens) <= 1 {
		return ""
	}
	return tokens[len(tokens)-1]
}
