package commands

import (
	"fmt"
	"os"
)

func GraphHelp() string {
	return `graph
Example, graph`
}

func Graph(args []string) {
	cmd := `git log --graph --oneline --branches --color --pretty=format:"%x1b[31m%h%x09%x1b[32m%d%x1b[0m%x20%s"`
	if out, err := RunCmd(cmd); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running graph command: ", err)
		os.Exit(1)
	} else {
		fmt.Fprintln(os.Stderr, out)
	}
}
