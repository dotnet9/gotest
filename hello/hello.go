package main

import (
	"fmt"

	"dotnet9.com/greetings"
)

func main() {
	message := greetings.Hello("Dotnet9")
	fmt.Println(message)
}