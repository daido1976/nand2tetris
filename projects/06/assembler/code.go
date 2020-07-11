package main

// CodeDest returns binary code of dest mnemonic
func CodeDest(mnemonic string) string {
	switch mnemonic {
	case "":
		return "000"
	case "M":
		return "001"
	case "D":
		return "010"
	case "MD":
		return "011"
	case "A":
		return "100"
	case "AM":
		return "101"
	case "AD":
		return "110"
	case "AMD":
		return "111"
	}
	// FIXME: error を返すようにした方がいいかも
	return ""
}
