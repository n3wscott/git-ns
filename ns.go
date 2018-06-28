package main

import (
	"./cmd"
	"os"
)

func main() {
	switch os.Args[1] {
	case "checkout":
		fallthrough
	case "branch":
		fallthrough
	case "co":
		cmd.Checkout(os.Args[2:])
	case "open":
		fallthrough
	case "o":
		cmd.Open(os.Args[2:])
	case "status":
		cmd.Status(os.Args[2:])
	case "sync":
		cmd.Sync(os.Args[2:])
	default:
		cmd.Help(os.Args[2:])
	}
}
