package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	asmFilePath := os.Args[len(os.Args)-1]
	rf, err := os.Open(asmFilePath)
	if err != nil {
		fmt.Println("error")
	}

	scanner := bufio.NewScanner(rf)
	parser := NewParser(scanner)

	hackFilePath := strings.ReplaceAll(asmFilePath, ".asm", ".hack")
	wf, err := os.Create(hackFilePath)
	if err != nil {
		fmt.Println("error")
	}

	// initialize
	symbolTable := NewSymbolTable()
	romAddress := 0

	// first path
	for parser.HasMoreCommands() {
		parser.Advance()

		switch parser.CommandType() {
		case A_COMMAND, C_COMMAND:
			romAddress++
		case L_COMMAND:
			symbolTable.AddEntry(parser.Symbol(), romAddress)
		}
	}

	// second path
	rf.Seek(0, 0)
	scanner = bufio.NewScanner(rf)
	parser = NewParser(scanner)
	ramAddress := 16

	for parser.HasMoreCommands() {
		parser.Advance()

		switch parser.CommandType() {
		case A_COMMAND:
			symbol := parser.Symbol()
			address, err := strconv.Atoi(symbol)
			if err == nil {
				// Xxx is number
				writeAddress(wf, address)
			} else {
				// Xxx is symbol
				if symbolTable.Contains(symbol) {
					// known symbol
					address := symbolTable.GetAddress(symbol)
					writeAddress(wf, address)
				} else {
					// new variables
					symbolTable.AddEntry(symbol, ramAddress)
					writeAddress(wf, ramAddress)
					ramAddress++
				}
			}
		case C_COMMAND:
			writeCInst(wf, parser)
		case L_COMMAND:
			// do nothing
		}
	}
}

func writeAddress(file *os.File, address int) {
	out := fmt.Sprintf("%016b", address)
	file.WriteString(out + "\n")
}

func writeCInst(file *os.File, parser *Parser) {
	out := "111" + CodeComp(parser.Comp()) + CodeDest(parser.Dest()) + CodeJump(parser.Jump())
	file.WriteString(out + "\n")
}
