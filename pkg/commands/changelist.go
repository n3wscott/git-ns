package commands

import (
	"fmt"
	"os"
	"strings"
)

func ChangeListHelp() string {
	return `cl
	Example, cl`
}

func ChangeList(args []string) {
	var diffFiles string
	var err error
	if diffFiles, err = RunCmd("git diff --name-only master", "quiet"); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git diff command: ", err)
		os.Exit(1)
	}

	files := strings.Split(diffFiles, "\n")

	for _, v := range files {
		v = strings.TrimSpace(v)
		if len(v) == 0 {
			continue
		}

		// file could have been deleted in the current diff.
		if !Exists(v) {
			fmt.Fprintln(os.Stderr, fmt.Sprintf("🗑 ", v))
		} else {
			fmt.Fprintln(os.Stderr, "✏️ ", v)
		}
	}
}
