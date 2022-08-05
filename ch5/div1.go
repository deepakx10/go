package main

import (
	"fmt"
	"os"
)

func divAndRem (numerator int, denominator int) (int, int, string) {
if denominator == 0 {
	return 0, 0, "nil" //error.New("Cannot devide by 0")
	}
return numerator / denominator, numerator % denominator, "nil"
}

func main() {
	result, reminder, err := divAndRem(8, 2)
	if err != "nil" {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, reminder)
}