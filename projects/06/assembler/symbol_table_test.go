package main

import (
	"testing"
)

func TestSymbol_Table(t *testing.T) {
	s := NewSymbolTable()

	if s.Contains("LOOP") {
		t.Fatalf("s must not contains LOOP")
	}

	s.AddEntry("LOOP", 16)
	if !s.Contains("LOOP") {
		t.Fatalf("s must contains LOOP")
	}

	actual := s.GetAddress("LOOP")
	if actual != 16 {
		t.Fatalf("Expecting LOOP address is 16, but %v", actual)
	}
}
