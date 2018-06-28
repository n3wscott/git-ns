package cmd

import (
	"fmt"
	"os"
)

func Checkout(args []string) {
	defaultBranch := fmt.Sprintf("branch-%s", RandString(6))
	branch := defaultBranch

	if len(args) > 0 && args[0] != "" {
		branch = args[0]
	}

	exists := false
	if _, err := RunCmd("git rev-parse --verify --quiet " + branch); err == nil {
		exists = true
	}

	flag := "-b "
	if exists {
		flag = ""
	}

	if _, err := RunCmd("git checkout " + flag + branch); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git command: ", err)
		os.Exit(1)
	}

	if !exists {
		if _, err := RunCmd("git push --set-upstream origin " + branch); err != nil {
			fmt.Fprintln(os.Stderr, "There was an error running git command: ", err)
			os.Exit(1)
		}
	}
}
