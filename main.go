package main

import (
	"DGFC/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hey %s! This is the REPL to test the compiler.\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
