package main

import (
	"testing"
)

type expectData struct {
	in  string
	out string
}

var expectsDest = []expectData{
	{"", "000"},
	{"M", "001"},
	{"D", "010"},
	{"MD", "011"},
	{"A", "100"},
	{"AM", "101"},
	{"AD", "110"},
	{"AMD", "111"},
	{"INVALID", ""},
}

func TestCode_Dest(t *testing.T) {
	for _, v := range expectsDest {
		out := CodeDest(v.in)
		if v.out != out {
			t.Errorf("CodeDest(%v) returns %v, but wants %v",
				v.in, out, v.out)
		}
	}
}
