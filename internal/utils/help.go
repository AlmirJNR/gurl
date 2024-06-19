package utils

import (
	"fmt"
	"os"
)

func ShowHelpMessage() {
	// todo: create help message
	fmt.Printf("Usage: %s <command> [<args>]", os.Args[0])
	fmt.Println()
}
