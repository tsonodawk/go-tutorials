package main

import (
	"fmt"

	"log"

	"example.com/greetings"
	// "rsc.io/quote"
)

func main() {
	// fmt.Println("Hello, World!")
	// fmt.Println(quote.Go())

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	// message, err := greetings.Hello("Gladys")
	// message, err := greetings.Hello("")
	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)

	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// Get a greeting message and print it.
	// fmt.Println(message)
	// If no error was returned, print the returned map of
	// messages to the console.
	fmt.Println(messages)

	for key := range messages {
		fmt.Printf("key: %v / message %v\n", key, messages[key])
	}

}
