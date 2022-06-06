package main

import "fmt"

import "rsc.io/quote"

func main() {
	fmt.Println(quote.Go())
}

func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
