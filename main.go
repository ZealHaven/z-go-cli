package main

import (
	"fmt"
	"z-go-cli/prompt"
)

func main() {
	name := prompt.StringPrompt("What is your name? ")
	fmt.Printf("Hello %s", name)
}
