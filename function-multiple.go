package main

import "fmt"

func fullName(f, l string, a uint) (string, string, uint) {
	return f, l, a
}

func main() {

	f, l, a := fullName("Elon", "musk", 29)

	fmt.Println(f, l, a)
	fmt.Println(fullName("Elon", "musk", 28))
}
