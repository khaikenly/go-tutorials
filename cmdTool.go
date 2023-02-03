package main

import (
	"fmt"
	"os"
)

func main() {
	// Print a welcome message
	fmt.Println("Welcome to your command line tool!")
	// Get the command line arguments
	args := os.Args[1:]
	// Check if any arguments were passed
	if len(args) == 0 {
		fmt.Println("No arguments were passed.")
		return
	}
	// Print the arguments
	fmt.Println("Arguments passed:")
	for _, arg := range args {
		fmt.Println(arg)
	}
}
