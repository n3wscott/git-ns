package cmd

import (
	"fmt"
	"os"
	"strings"
)

func OpenHelp() string {
	return `open [remote]
	remote: default=upstream
Example, open upstream`
}

func Open(args []string) {
	defaultRemote := "origin"
	remote := defaultRemote

	if len(args) > 0 && args[0] != "" {
		remote = args[0]
	}

	s := fmt.Sprintf("remote.%s.url", remote)

	cmdOut, err := RunCmd("git config --get " + s)
	if err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git command: ", err)
		os.Exit(1)
	}

	giturl := string(cmdOut)
	giturl = strings.Replace(giturl, "\n", "", 1)

	if strings.Contains(giturl, "git@") {
		// This is an ssh url.
		giturl = strings.Replace(giturl, "git@", "https://", 1)
		giturl = strings.Replace(giturl, ".git", "/", 1)
		giturl = strings.Replace(giturl, "com:", "com/", 1)
	} else {
		// This is an http url.
		fmt.Println("http")
	}

	fmt.Println(giturl)

	if _, err := RunCmd("open " + giturl); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running open command: ", err)
		os.Exit(1)
	}
}
