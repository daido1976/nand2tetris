package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	f, err := os.Open("projects/06/max/Max.asm")
	if err != nil {
		fmt.Println("error")
	}

	s := bufio.NewScanner(f)
	p := newParser(s)

	for p.hasMoreCommands() {
		p.advance()
		if p.commandType() == A_COMMAND || p.commandType() == L_COMMAND {
			fmt.Println(p.commandType(), p.currentCommand, "->", p.symbol())
		} else {
			fmt.Println(p.commandType(), p.currentCommand)
		}
	}
}

// test
func fileOut() {
	file, err := os.Open("./projects/06/add/Add.asm")
	if err != nil {
		fmt.Println("error")
	}
	defer file.Close()
	fmt.Printf("name: %s\n", file.Name())

	s := bufio.NewScanner(file)

	// return to line 10
	for i := 1; s.Scan() && i < 10; i++ {
		line := s.Text()
		fmt.Println(i, line)
	}
}

func newParser(s *bufio.Scanner) *Parser {
	return &Parser{s, ""}
}

func (p *Parser) hasMoreCommands() bool {
	return p.scanner.Scan()
}

func (p *Parser) advance() {
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

func (p *Parser) commandType() Command {
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

func (p *Parser) symbol() string {
	if p.commandType() == A_COMMAND {
		// A_COMMAND: @Xxx -> Xxx
		return strings.TrimLeft(p.currentCommand, "@")
	}
	// L_COMMAND: (Xxx) -> Xxx
	return strings.TrimRight(strings.TrimLeft(p.currentCommand, "("), ")")
}
