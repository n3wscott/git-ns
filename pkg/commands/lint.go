package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/client9/misspell"
)

func LintHelp() string {
	return `lint
	Example, lint`
}

func Lint(args []string) {
	var diffFiles string
	var err error
	if diffFiles, err = RunCmd("git diff --name-only master"); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git diff command: ", err)
		os.Exit(1)
	}

	files := strings.Split(diffFiles, "\n")

	for _, v := range files {
		v = strings.TrimSpace(v)
		if len(v) == 0 {
			continue
		}

		switch {
		case strings.HasSuffix(v, ".go"):
			lintGo(v)
		case strings.HasSuffix(v, ".md"):
			lintMarkdown(v)
		default:
			fmt.Fprintln(os.Stderr, fmt.Sprintf("⚠️  ignored: %s", v))
		}
	}

}

func lintGo(file string) {
	if changed, err := RunCmd(fmt.Sprintf("go fmt %s", file)); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git fmt command: ", err)
		os.Exit(1)
	} else if len(changed) > 0 {
		fmt.Fprintln(os.Stderr, "➡️ ", file)
	} else {
		fmt.Fprintln(os.Stderr, "✅", file)
	}

}

func lintMarkdown(file string) {

}

var replacer *misspell.Replacer

func spelling(file string) {
	if replacer == nil {
		r := &misspell.Replacer{
			Replacements: misspell.DictMain,
			Debug:        false,
		}
		r.Compile()
	}
}