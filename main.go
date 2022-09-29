package main

import (
	"bufio"
	"os"
	"fmt"
	"io"
	"strings"
)

var cmdFunc func(w io.Writer, args []string) (exit bool)

func main() {
	s := bufio.NewScanner(os.Stdin)
	w := os.Stdout
	fmt.Println("hello to amg terminal")
	for {
		s.Scan()
		args := strings.Split(string(s.Bytes()), " ")
		cmd := args[0]
		args = args[1:]

		switch cmd {
		case "exit":
			cmdFunc = exitCmd
		}

		if cmdFunc == nil {
			fmt.Fprintf(w, "%q not found\n", cmd)
			continue
		}

		if cmdFunc(w, args) {
			return
		}
		
	}
}

func exitCmd(w io.Writer, args []string) bool {
	fmt.Fprintf(w, "so long and thanks for all the fish\n")
	return true
}