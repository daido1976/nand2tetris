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
				out := fmt.Sprintf("%016b", address)
				wf.WriteString(out + "\n")
			} else {
				// Xxx is symbol
				if symbolTable.Contains(symbol) {
					// known symbol
					address := symbolTable.GetAddress(symbol)
					out := fmt.Sprintf("%016b", address)
					wf.WriteString(out + "\n")
				} else {
					// new variables
					symbolTable.AddEntry(symbol, ramAddress)
					out := fmt.Sprintf("%016b", ramAddress)
					wf.WriteString(out + "\n")
					ramAddress++
				}
			}
		case C_COMMAND:
			out := "111" + CodeComp(parser.Comp()) + CodeDest(parser.Dest()) + CodeJump(parser.Jump())
			wf.WriteString(out + "\n")
		case L_COMMAND:
			// do nothing
		}
	}
}
