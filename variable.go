package main

import "fmt"

// say := "hello"  <-- failed calling this variable
var say string = "hello"

func main() {

	name := "yuuki"
	var age = 20
	var book int = 50
	bol := true

	fmt.Println(say)
	fmt.Println("her name is", name)
	fmt.Printf("she is %v years old \n", age)
	fmt.Printf("she had %v books \n", book)
	fmt.Println("boolean", bol)
}
