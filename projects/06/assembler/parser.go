package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileOut()
}

func fileOut() {
	file, err := os.Open("./projects/06/add/Add.asm")
	if err != nil {
		fmt.Println("error")
	}
	defer file.Close()
	fmt.Printf("name: %s\n", file.Name())

	s := bufio.NewScanner(file)

	// return to line 10
	for i := 1; s.Scan() && i < 10; i++ {
		line := s.Text()
		fmt.Println(i, line)
	}
}
