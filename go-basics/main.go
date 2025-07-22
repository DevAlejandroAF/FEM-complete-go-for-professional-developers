package main

import (
	"example.com/basics/imports"
	"fmt"
)

func main() {
	myTicket := imports.Ticket{
		ID:    123,
		Event: "theater",
	}

	myTicket.ChangeEvent("new event")

	fmt.Println("Hello world")
}
