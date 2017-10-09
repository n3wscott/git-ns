package main

import (
	"./cmd"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, world! Args:")
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	cmd.Status()
}
