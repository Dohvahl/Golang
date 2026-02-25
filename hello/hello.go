package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	// Get a greeting message and prirnt it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
