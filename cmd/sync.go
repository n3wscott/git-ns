package cmd

import (
	"fmt"
	"os"
)

func Sync(args []string) {
	if HasPendingChanges() {
		fmt.Fprintln(os.Stderr, "There are pending changes in the current branch.")
		os.Exit(1)
	}

	cmdOut, err := RunCmd("git status")
	if err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git command: ", err)
		os.Exit(1)
	}

	fmt.Println(string(cmdOut))
	fmt.Println("TODO")
}
