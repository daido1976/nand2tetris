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
	f, err := os.Open(asmFilePath)
	if err != nil {
		fmt.Println("error")
	}

	s := bufio.NewScanner(f)
	p := NewParser(s)

	hackFilePath := strings.ReplaceAll(asmFilePath, ".asm", ".hack")
	wf, err := os.Create(hackFilePath)
	if err != nil {
		fmt.Println("error")
	}

	// initialize
	st := NewSymbolTable()
	romAddress := 0

	// first path
	for p.HasMoreCommands() {
		p.Advance()

		switch p.CommandType() {
		case A_COMMAND, C_COMMAND:
			romAddress++
		case L_COMMAND:
			st.AddEntry(p.Symbol(), romAddress)
		}
	}

	// second path
	f.Seek(0, 0)
	s = bufio.NewScanner(f)
	p = NewParser(s)
	ramAddress := 16

	for p.HasMoreCommands() {
		p.Advance()

		switch p.CommandType() {
		case A_COMMAND:
			symbol := p.Symbol()
			address, err := strconv.Atoi(symbol)
			if err == nil {
				// Xxx is number
				out := fmt.Sprintf("%016b", address)
				wf.WriteString(out + "\n")
			} else {
				// Xxx is symbol
				if st.Contains(symbol) {
					// known symbol
					address := st.GetAddress(symbol)
					out := fmt.Sprintf("%016b", address)
					wf.WriteString(out + "\n")
				} else {
					// new variables
					st.AddEntry(symbol, ramAddress)
					out := fmt.Sprintf("%016b", ramAddress)
					wf.WriteString(out + "\n")
					ramAddress++
				}
			}
		case C_COMMAND:
			out := "111" + CodeComp(p.Comp()) + CodeDest(p.Dest()) + CodeJump(p.Jump())
			wf.WriteString(out + "\n")
		case L_COMMAND:
			// do nothing
		}
	}
}
