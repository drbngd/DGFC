package main

import (
	"DGFC/cmd/repl"
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

// test to see if .idea is being ignored
