package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

func main() {
	fmt.Println(math.Pi)
	fmt.Println(math.Ceil(78.68))
	fmt.Println(math.Floor(78.48))
	fmt.Println(rand.Intn(50))

	fmt.Println(strings.Count("hello", "l"))
	fmt.Println(strings.ReplaceAll("I am able to build my own start up", " am", "'m"))
}
