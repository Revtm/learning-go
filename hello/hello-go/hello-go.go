package main

import "rsc.io/quote"
import (
	"fmt"
	"example.com/greetings"
)

func main() {
	fmt.Println(quote.Go())
	message := greetings.Hello("Revtm")
	fmt.Println(message)
}
