package main

import "fmt"

func add(vals ...int) int {
	result := 0
	for _, n := range vals {
		result = result+n
	}
	return result
}

func main() {
	fmt.Println(add(10, 20, 30))
}