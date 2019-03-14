package commands

import (
	"fmt"
	"os"
	"strings"
)

func CloneHelp() string {
	return `clone [upstream git url]
	Clones and then sets upstream.
	Example: clone https://github.com/n3wscott/git-tools`
}

func Clone(args []string) {
	var err error
	var forkOrg string
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "upstream url required.\n", CloneHelp())
		os.Exit(1)
	}
	upstream := strings.TrimSpace(args[0])

	if forkOrg, err = RunCmd("git config --global --get ns.fork.org"); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git command: ", err)
		os.Exit(1)
	}
	forkOrg = strings.TrimSpace(forkOrg)

	parts := strings.Split(upstream, "/")

	if len(parts) != 5 || parts[0] != "https:" {
		fmt.Fprintln(os.Stderr, "Not sure how to handle given url: ", upstream)
		os.Exit(1)
	}
	host := parts[2]
	upstreamOrg := parts[3]
	repo := parts[4]

	var pwd string
	if pwd, err = RunCmd("pwd"); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running pwd command: ", err)
		os.Exit(1)
	}
	pwd = strings.TrimSpace(pwd)
	pwdParts := strings.SplitAfter(pwd, host)
	hostDir := pwdParts[0]

	if !strings.Contains(pwd, host) {
		fmt.Fprintln(os.Stderr, "[error] Forcing directory structure to be golang style. Want", host, "in current path.")
		os.Exit(1)
	}

	var orgLs string
	if orgLs, err = RunCmdAt("ls", hostDir); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running ls command: ", err)
		os.Exit(1)
	}

	if !strings.Contains(orgLs, upstreamOrg+"\n") {
		if _, err = RunCmdAt(fmt.Sprintf("mkdir %s", upstreamOrg), hostDir); err != nil {
			fmt.Fprintln(os.Stderr, "There was an error running mkdir command: ", err)
			os.Exit(1)
		}
	}

	orgDir := fmt.Sprintf("%s/%s", hostDir, upstreamOrg)

	fork := fmt.Sprintf("git@%s:%s/%s.git", host, forkOrg, repo)
	upstream = upstream + ".git"

	fmt.Fprintln(os.Stderr, "upstream:", upstream)

	fmt.Fprintln(os.Stderr, "fork:", fork)

	if _, err = RunCmdAt(fmt.Sprintf("git clone %s", fork), orgDir); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git clone command: ", err)
		os.Exit(1)
	}

	repoDir := fmt.Sprintf("%s/%s", orgDir, repo)

	if _, err = RunCmdAt(fmt.Sprintf("git remote add upstream %s", upstream), repoDir); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git remote add upstream command: ", err)
		os.Exit(1)
	}

	if _, err = RunCmdAt("git remote set-url --push upstream no_push", repoDir); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git remote set-url command: ", err)
		os.Exit(1)
	}

	if _, err = RunCmdAt("git ns sync", repoDir); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git ns sync command: ", err)
		os.Exit(1)
	}

}
