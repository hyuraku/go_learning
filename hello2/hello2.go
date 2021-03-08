package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Tom")
	fmt.Println(message)
}
