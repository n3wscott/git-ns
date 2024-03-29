package commands

import (
	"fmt"
	"os"
)

func Sync(args []string) {
	if HasPendingChanges() {
		fmt.Fprintln(os.Stderr, "There are pending changes in the current branch.")
		os.Exit(1)
	}

	if _, err := RunCmd("git fetch upstream"); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git command: ", err)
		os.Exit(1)
	}

	branch := DefaultBranch("upstream")

	// TODO: this assumes master, that is wrong. git remote show upstream | grep "HEAD branch" | cut -d ":" -f 2
	if _, err := RunCmd("git rebase upstream/" + branch); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git command: ", err)
		os.Exit(1)
	}

}
