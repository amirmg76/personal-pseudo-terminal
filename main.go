package main

import (
	"bufio"
	"os"
	"fmt"
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
		message := string(s.Bytes())
		if message == "exit" {
			return
		}
	}
}