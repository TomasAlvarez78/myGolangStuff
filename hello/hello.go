package main

import (
    "fmt"
	"log"

    "example.com/greetings"
)

func main() {
    // Get a greeting message and print it.
	log.SetPrefix("Greetings: ")
	log.SetFlags(0)

	names := []string{
		"Elira",
		"Emma",
		"Pauli",
	}
	
    messages,err := greetings.Tasuketes(names)

	if err != nil {
		log.Fatal(err)
	}

    fmt.Println(messages)
}
