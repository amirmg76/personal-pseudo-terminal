package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	w := os.Stdout
	fmt.Println("hello to amg terminal")
	for {
		s.Scan()
		fmt.Print("you wrote ")
		w.Write(s.Bytes())
		fmt.Println()
		args := strings.Split(string(s.Bytes()), " ")
		cmd := args[0]
		args = args[1:]
		if cmd == "exit" {
			return
		}
		
	}
}