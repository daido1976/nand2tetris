package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// open file for reading with last command line args
	asmFilePath := os.Args[len(os.Args)-1]
	rf, err := os.Open(asmFilePath)
	if err != nil {
		fmt.Println("Error! The file can not open")
	}

	// initialize scanner & parser
	scanner := bufio.NewScanner(rf)
	parser := NewParser(scanner)

	// first pass to scan symbol
	symbolTable := NewSymbolTable()
	romAddress := 0
	for parser.HasMoreCommands() {
		parser.Advance()

		switch parser.CommandType() {
		case A_COMMAND, C_COMMAND:
			romAddress++
		case L_COMMAND:
			symbolTable.AddEntry(parser.Symbol(), romAddress)
		}
	}

	// create file for writing
	hackFilePath := strings.ReplaceAll(asmFilePath, ".asm", ".hack")
	wf, err := os.Create(hackFilePath)
	if err != nil {
		fmt.Println("Error! The file can not create")
	}

	// reset scanner & parser
	rf.Seek(0, 0)
	scanner = bufio.NewScanner(rf)
	parser = NewParser(scanner)

	// second pass to assemble
	ramAddress := 16
	for parser.HasMoreCommands() {
		parser.Advance()

		switch parser.CommandType() {
		case A_COMMAND:
			symbol := parser.Symbol()
			address, err := strconv.Atoi(symbol)
			if err == nil {
				// Xxx is number
				writeAInst(wf, address)
			} else {
				// Xxx is symbol
				if symbolTable.Contains(symbol) {
					// a known symbol is converted to an address
					address := symbolTable.GetAddress(symbol)
					writeAInst(wf, address)
				} else {
					// an unknown symbol is treated as a new variable
					symbolTable.AddEntry(symbol, ramAddress)
					writeAInst(wf, ramAddress)
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

// writeAInst writes address instruction.
// mnemonic: @value
// binary: 0vvv vvvv vvvv vvvv
func writeAInst(file *os.File, address int) {
	out := fmt.Sprintf("%016b", address)
	file.WriteString(out + "\n")
}

// writeCInst writes comtute instruction.
// mnemonic: dest=comp;jump
// binary: 111a cccc ccdd djjj
func writeCInst(file *os.File, parser *Parser) {
	out := "111" + CodeComp(parser.Comp()) + CodeDest(parser.Dest()) + CodeJump(parser.Jump())
	file.WriteString(out + "\n")
}
