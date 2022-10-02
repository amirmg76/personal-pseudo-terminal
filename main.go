package main

import (
	"bufio"
	"math/rand"
	"os"
	"fmt"
	"io"
	"strings"
)

type cmd struct {
	Name string
	Help string
	Action func(w io.Writer, args []string) bool
}

var cmdFunc func(w io.Writer, args []string) (exit bool)

func main() {
	s := bufio.NewScanner(os.Stdin)
	w := os.Stdout
	fmt.Println("hello to amg terminal")
	cmds := []cmd{}
	help := cmd{
		Name: "help",
		Help: "Shows available commands",
		Action: func(w io.Writer, args []string) bool {
			fmt.Fprintln(w, "Available commands:")
			for _, c := range cmds {
				fmt.Fprintf(w, " - %-15s %s\n", c.Name, c.Help)
			}
			return false
		},
	}

	cmds = append(cmds, help)
	exit := cmd{
		Name: "exit",
		Help: "let you exit the terminal",
		Action: func(w io.Writer, args []string) bool {
			fmt.Fprintf(w, "so long and thanks for all the fish")
			return true
		},
	}
	cmds = append(cmds, exit)

	shuffle := cmd{
		Name: "shuffle",
		Help: "shuffle your args",
		Action: func(w io.Writer, args []string) bool {
			rand.Shuffle(len(args), func(i, j int){
				args[i], args[j] = args[j], args[i]
			})
			for i := range args {
				fmt.Fprintf(w, "%s ", args[i])
			}
			fmt.Fprintln(w)
			return false
		},
	}
	cmds = append(cmds, shuffle)

	echo := cmd{
		Name: "echo",
		Help: "it echos what you said",
		Action: func(w io.Writer, args []string) bool {
			for i := range args {
				fmt.Fprintf(w, "%s ", args[i])
			}
			fmt.Fprintln(w)
			return false
		},
	}
	cmds = append(cmds, echo)
	
	print := cmd{
		Name: "print",
		Help: "it print file's content",
		Action: func(w io.Writer, args []string) bool {
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
		},
	}
	cmds = append(cmds, print)
	
	for {
		s.Scan()
		args := strings.Split(string(s.Bytes()), " ")
		idx := -1
		for i := range cmds {
			if !cmds[i].Match(args[0]) {
				continue
			}
			idx = i
			break
		}
		if idx == -1 {
			fmt.Fprintf(w, "%q not found. Use `help` for available commands\n", args[0])
			continue
		}
		if cmds[idx].Run(w, args[1:]) {
			fmt.Fprintln(w)
			return
		}
	}
}

func (c cmd) Match (s string) bool {
	return c.Name == s
}

func (c cmd) Run(w io.Writer, args []string) bool {
	return c.Action(w, args)
}
