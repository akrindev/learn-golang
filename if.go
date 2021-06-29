package main

import "fmt"

func main() {

	if 1 < 2 {
		fmt.Println("yep 1 < 2")
	}

	// num := 1
	// num := 2
	num := 4

	if num == 1 {
		fmt.Println(num, "is equal")
	} else if num == 4 {
		fmt.Println(num, "is 4")
	} else {
		fmt.Println("nope")
	}

	fmt.Println()
}
