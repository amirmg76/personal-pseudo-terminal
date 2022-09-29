package main

import (
	"bufio"
	"math/rand"
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
		case "shuffle":
			cmdFunc = shuffleCmd
		case "echo":
			cmdFunc = echoFunc
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

func shuffleCmd(w io.Writer, args []string) bool {
	rand.Shuffle(len(args), func(i, j int){
		args[i], args[j] = args[j], args[i]
	})
	for i := range args {
		fmt.Fprintf(w, "%s ", args[i])
	}
	fmt.Fprintln(w)
	return false
}

func echoCmd(w io.Writer, args []string) bool {
	for i := range args {
		fmt.Fprintf(w, "%s ", args[i])
	}
	fmt.Fprintln(w)
	return false
}