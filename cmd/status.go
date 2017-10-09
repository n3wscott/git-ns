package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func Status() {
	var (
		cmdOut []byte
		err    error
	)
	cmdArgs := []string{"status"}
	if cmdOut, err = exec.Command("git", cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git command: ", err)
		os.Exit(1)
	}

	fmt.Println(string(cmdOut))
}
