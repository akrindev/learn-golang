package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()

	// fmt.Println(t)
	// name := "klein"
	name := "yui"

	switch name {
	case "klein":
		fmt.Println("his name is klein")
	case "yui":
		fmt.Println("oh its yui")
	default:
		fmt.Println("Whats your name?")
	}

	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good Evening")
	}

	fmt.Println("thanks")
}
