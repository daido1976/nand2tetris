package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("projects/06/max/Max.asm")
	if err != nil {
		fmt.Println("error")
	}

	s := bufio.NewScanner(f)
	p := NewParser(s)

	for p.HasMoreCommands() {
		p.Advance()
		if p.CommandType() == A_COMMAND || p.CommandType() == L_COMMAND {
			fmt.Println(p.CommandType(), p.currentCommand, "->", p.Symbol())
		} else if p.CommandType() == C_COMMAND {
			fmt.Println(p.CommandType(), p.currentCommand)
			fmt.Println("dest:"+p.Dest(), "comp:"+p.Comp(), "jump:"+p.Jump())
		} else {
			fmt.Println(p.CommandType(), p.currentCommand)
		}
	}
}
