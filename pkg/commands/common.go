package commands

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
)

func Cmd(cmdLine string, opts ...string) *exec.Cmd {
	if len(opts) == 0 {
		fmt.Fprintln(os.Stderr, cmdLine)
	}
	cmdSplit := strings.Split(cmdLine, " ")
	cmd := cmdSplit[0]
	args := cmdSplit[1:]

	return exec.Command(cmd, args...)
}

func RunCmd(cmdLine string, opts ...string) (string, error) {
	cmd := Cmd(cmdLine, opts...)

	cmdOut, err := cmd.Output()
	return string(cmdOut), err
}

func RunCmdAt(cmdLine, dir string, opts ...string) (string, error) {
	cmd := Cmd(cmdLine, opts...)
	cmd.Dir = dir

	cmdOut, err := cmd.Output()
	return string(cmdOut), err
}

func Exists(file string) bool {
	if _, err := RunCmd(fmt.Sprintf("test -f %s", file), "quiet"); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running test command: ", err)
		return false
	} else {
		return true
	}
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
