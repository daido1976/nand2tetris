package main

import (
	"bufio"
	"strings"
	"testing"
)

type expectParserData struct {
	commandType Command
	symbol      string
	dest        string
	comp        string
	jump        string
}

var asmCode = `// comment line
	@i
	M=1 // i=1
	@sum
	M=0 // sum=0

(LOOP)
	@i
	AMD=D-A;JMP
	0;JMP`

var expects = []expectParserData{
	{N_COMMAND, "", "", "", ""},
	{A_COMMAND, "i", "", "", ""},
	{C_COMMAND, "", "M", "1", ""},
	{A_COMMAND, "sum", "", "", ""},
	{C_COMMAND, "", "M", "0", ""},
	{N_COMMAND, "", "", "", ""},
	{L_COMMAND, "LOOP", "", "", ""},
	{A_COMMAND, "i", "", "", ""},
	{C_COMMAND, "", "AMD", "D-A", "JMP"},
	{C_COMMAND, "", "", "0", "JMP"},
}

func TestParser(t *testing.T) {
	buf := strings.NewReader(asmCode)
	scanner := bufio.NewScanner(buf)
	parser := NewParser(scanner)

	for i := range expects {
		if !parser.HasMoreCommands() {
			t.Errorf("unexpected finish at line: %d", i+1)
		}

		parser.Advance()
		expect := &expects[i]
		actualType := parser.CommandType()
		if expect.commandType != actualType {
			t.Fatalf("%v: expect %v, but %v", parser.currentCommand, expect.commandType, actualType)
		}

		switch actualType {
		case A_COMMAND, L_COMMAND:
			actualSymbol := parser.Symbol()
			if expect.symbol != actualSymbol {
				t.Fatalf("expect symbol %v, but %v", expect.symbol, actualSymbol)
			}
		case C_COMMAND:
			aDest := parser.Dest()
			aComp := parser.Comp()
			aJump := parser.Jump()
			if expect.dest != aDest || expect.comp != aComp || expect.jump != aJump {
				t.Fatalf("dest=comp;jump: %v=%v;%v want %v=%v;%v",
					aDest, aComp, aJump, expect.dest, expect.comp, expect.jump)
			}
		}
	}
}
