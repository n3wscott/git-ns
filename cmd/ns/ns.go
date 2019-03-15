package main

import (
	cmd "github.com/n3wscott/git-tools/pkg/commands"
	"os"
)

func main() {
	arg := ""
	if len(os.Args) >= 2 {
		arg = os.Args[1]
	}
	switch arg {
	case "clone":
		cmd.Clone(os.Args[2:])
	case "checkout", "branch", "co":
		cmd.Checkout(os.Args[2:])
	case "graph", "tree":
		cmd.Graph(os.Args[2:])
	case "lint":
		cmd.Lint(os.Args[2:])
	case "open", "o":
		cmd.Open(os.Args[2:])
	case "status":
		cmd.Status(os.Args[2:])
	case "sync":
		cmd.Sync(os.Args[2:])
	default:
		cmd.Help()
	}
}
