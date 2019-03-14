package commands

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
)

func RunCmd(cmdLine string) (string, error) {
	fmt.Fprintln(os.Stderr, cmdLine)
	cmdSplit := strings.Split(cmdLine, " ")
	cmd := cmdSplit[0]
	args := cmdSplit[1:]

	cmdOut, err := exec.Command(cmd, args...).Output()
	return string(cmdOut), err
}

func HasPendingChanges() bool {
	cmdOut, err := RunCmd("git diff-index --name-only HEAD --")
	if err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git command: ", err)
		os.Exit(1)
	}

	if len(cmdOut) > 0 {
		return true
	}
	return false
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
