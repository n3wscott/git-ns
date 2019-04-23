package commands

import (
	"fmt"
)

func Help() {
	fmt.Println("Use git ns sync or whatever")
	fmt.Println(ChangeListHelp())
	fmt.Println(CheckoutHelp())
	fmt.Println(OpenHelp())
	fmt.Println(GraphHelp())

	fmt.Println(CloneHelp())
	fmt.Println(LintHelp())
}
