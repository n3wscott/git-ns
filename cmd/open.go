package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func OpenHelp() string {
	return `open [remote]
	remote: default=upstream
Example, open upstream`
}

func Open(args []string) {
	var (
		cmdOut []byte
		err    error
	)

	defaultRemote := "origin"

	remote := defaultRemote

	if len(args) > 0 && args[0] != "" {
		remote = args[0]
	}

	//giturl=$(git config --get "remote.${remote}.url")
	s := fmt.Sprintf("remote.%s.url", remote)
	cmdArgs := []string{"config", "--get", s}
	if cmdOut, err = exec.Command("git", cmdArgs...).Output(); err != nil {
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

	if cmdOut, err = exec.Command("open", giturl).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running open command: ", err)
		os.Exit(1)
	}
}
