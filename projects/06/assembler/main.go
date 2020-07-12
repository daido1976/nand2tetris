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
	for p.HasMoreCommands() {
		p.Advance()
		if p.CommandType() == A_COMMAND || p.CommandType() == L_COMMAND {
			fmt.Println("Debug:", p.CommandType(), p.currentCommand, "->", p.Symbol())

			i, _ := strconv.Atoi(p.Symbol())
			out := fmt.Sprintf("%016b", i)
			// debug result
			fmt.Println(out)
			wf.WriteString(out + "\n")
		} else if p.CommandType() == C_COMMAND {
			fmt.Println("Debug:", p.CommandType(), p.currentCommand)
			fmt.Println("Debug:", "dest:"+p.Dest(), "comp:"+p.Comp(), "jump:"+p.Jump())

			out := "111" + CodeComp(p.Comp()) + CodeDest(p.Dest()) + CodeJump(p.Jump())
			// debug result
			fmt.Println(out)
			wf.WriteString(out + "\n")
		}
	}
}
