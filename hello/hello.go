package main

import (
	"fmt"
	"log"

	"dotnet9.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Dotnet9", "C#", "Go"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	} 
	
	fmt.Println(messages)
}