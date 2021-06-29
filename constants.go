package main

import "fmt"

// constants cant be declarad using := syntax
// const age := 20 <-- this will error
const Pi = 3.14

func main() {
	const world = "World"

	fmt.Println("Hellow", world)
	fmt.Println("Pi:", Pi)

	const yes = true
	fmt.Println("Yes?", yes)
}
