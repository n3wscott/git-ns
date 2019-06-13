package commands

import (
	"fmt"
	"os"
	"path/filepath"
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
			continue
		}

		switch {
		case strings.HasPrefix(v, "vendor/"):
			continue
		case strings.HasSuffix(v, ".go"):
			lintGo(v)
		case strings.HasSuffix(v, ".md"):
			lintMarkdown(v)
		default:
			fmt.Fprintln(os.Stderr, fmt.Sprintf("⚠️  ignored: %s", v))
		}
		//checkSpelling(v) <<- does not seem to work
	}

}

func lintGo(file string) {
	lintGoFmt(file)
	lintGoImports(file)
}

func lintGoFmt(file string) {
	if changed, err := RunCmd(fmt.Sprintf("gofmt -s -w -l %s", file)); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git fmt command: ", err)
		os.Exit(1)
	} else if len(changed) > 0 {
		fmt.Fprintln(os.Stderr, "➡️ fmt", file)
	} else {
		fmt.Fprintln(os.Stderr, "✅ fmt", file)
	}
}

func lintGoImports(file string) {
	if changed, err := RunCmd(fmt.Sprintf("goimports -w -l %s", file)); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git fmt command: ", err)
		os.Exit(1)
	} else if len(changed) > 0 {
		fmt.Fprintln(os.Stderr, "➡️ ", file)
	} else {
		fmt.Fprintln(os.Stderr, "✅ imports", file)
	}
}

func lintMarkdown(file string) {
	if changed, err := RunCmd(fmt.Sprintf("prettier --write --list-different --no-color --prose-wrap=always %s", file)); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running prettier command: ", err)
		os.Exit(1)
	} else if len(changed) > 0 {
		fmt.Fprintln(os.Stderr, "➡️ ", file)
	} else {
		fmt.Fprintln(os.Stderr, "✅", file)
	}
}

var r *misspell.Replacer

func checkSpelling(file string) {
	if r == nil {
		r = &misspell.Replacer{
			Replacements: misspell.DictMain,
			Debug:        false,
		}
		r.Compile()
	}

	if diff, err := RunCmd(fmt.Sprintf("git diff master %s", file)); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git diff command: ", err)
		return
	} else {
		for _, line := range strings.Split(diff, "\n") {

			if strings.HasPrefix(line, "++") {
				// this is a new file,
				continue
			} else if strings.HasPrefix(line, "+") {
				orig := line[1:]
				updated := orig // Default to skip comment when to files match.

				if filepath.Ext(file) == ".go" {
					fmt.Fprintln(os.Stderr, "Looking at go ", orig)
					updated, _ = r.ReplaceGo(orig)
				} else if filepath.Ext(file) == ".md" {
					fmt.Fprintln(os.Stderr, "Looking at md", orig)
					updated, _ = r.Replace(orig)
				}
				if updated != orig {
					// handelde
					fmt.Fprintln(os.Stderr, "->>  ", file)
				}
			}
		}
	}
}
