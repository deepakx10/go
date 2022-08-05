package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	switch n := rand.Intn; {
	case n == 0:
		fmt.Println(n, "is too low")
	case n > 5:
		fmt.Println(n, "is big")
	default:
		fmt.Println("good number")
	}
}