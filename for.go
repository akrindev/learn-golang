package main

import "fmt"

func main() {

	// go has only one for loop

	// for init; condition; post { ... }
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	first, end := 1, 5

	for first < end { // same as while looping
		fmt.Println("first is still less than end")
		first++ // this will cause infinite loop if first is always less than end
	}

	fmt.Println("uh ended for looping")
}
