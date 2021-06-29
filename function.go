package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func sayHello(s string) string {
	return "Hellow " + s
}

func main() {
	fmt.Println(add(4, 5))
	fmt.Println(sub(4, 5))
	fmt.Println(sayHello("yuuki"))
}
