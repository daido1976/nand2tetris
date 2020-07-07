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

func main() {
	f, err := os.Open("./projects/06/add/Add.asm")
	if err != nil {
		fmt.Println("error")
	}

	s := bufio.NewScanner(f)
	p := newParser(s)

	for p.hasMoreCommands() {
		p.advance()
		fmt.Println(p.currentCommand)
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
