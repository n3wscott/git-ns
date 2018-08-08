package main

import (
	"./cmd"
	"os"
)

func main() {
	arg := ""
	if len(os.Args) >= 2 {
		arg = os.Args[1]
	}
	switch arg {
	case "checkout", "branch", "co":
		cmd.Checkout(os.Args[2:])
	case "graph", "tree":
		cmd.Graph(os.Args[2:])
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
