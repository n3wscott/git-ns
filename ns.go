package main

import (
	"./cmd"
	"os"
)

func main() {
	switch os.Args[1] {
	case "checkout":
	case "branch":
	case "co":
		cmd.Checkout(os.Args[2:])
	case "open":
	case "o":
		cmd.Open(os.Args[2:])
	case "status":
		cmd.Status(os.Args[2:])
	case "sync":
	case "s":
		cmd.Sync(os.Args[2:])
	default:
		cmd.Help(os.Args[2:])
	}
}
