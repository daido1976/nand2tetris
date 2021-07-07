package main

// SymbolTable records a table of symbols and addresses.
type SymbolTable map[string]int

// NewSymbolTable returns new SymbolTable.
func NewSymbolTable() SymbolTable {
	return SymbolTable{
		"SP":     0,
		"LCL":    1,
		"ARG":    2,
		"THIS":   3,
		"THAT":   4,
		"R0":     0,
		"R1":     1,
		"R2":     2,
		"R3":     3,
		"R4":     4,
		"R5":     5,
		"R6":     6,
		"R7":     7,
		"R8":     8,
		"R9":     9,
		"R10":    10,
		"R11":    11,
		"R12":    12,
		"R13":    13,
		"R14":    14,
		"R15":    15,
		"SCREEN": 16384,
		"KBD":    24576,
	}
}

// AddEntry adds a symbol and address pair to the table.
func (s SymbolTable) AddEntry(symbol string, address int) {
	s[symbol] = address
}

// Contains returns whether or not the symbol table contains a given symbol.
func (s SymbolTable) Contains(symbol string) bool {
	_, ok := s[symbol]
	return ok
}

// GetAddress returns the address associated with the symbol.
func (s SymbolTable) GetAddress(symbol string) int {
	address := s[symbol]
	return address
}
