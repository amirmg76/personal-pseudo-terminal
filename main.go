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
			cmdFunc = echoCmd
		case "print":
		    cmdFunc = printCmd
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

func printCmd(w io.Writer, args []string) bool {
	if len(args) != 1 {
		fmt.Fprintln(w, "please specify only one file")
		return false
	}
	file, err := os.Open(args[0])
	if err != nil {
		fmt.Fprintf(w, "Cannot open %s\nError message: %s\n", args[0], err)
		fmt.Fprintln(w)
		return false
	} 
	defer file.Close()
	if _, err := io.Copy(w, file); err != nil {
		fmt.Fprintln(w, "Cannot print %s\nError message: %s\n", args[0], err)
		return false
	}
	fmt.Fprintln(w)
	return false
}