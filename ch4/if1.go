package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := rand.Intn(10)
	if n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too high")
	} else {
		fmt.Println(n)
		fmt.Println("That's about rght")
	}
}