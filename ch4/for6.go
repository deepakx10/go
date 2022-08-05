package main

import (
	"fmt"
)

func main() {
	samples := []string{"hello", "apple"}
	for _, sample := range samples {
		for k, v := range sample {
			fmt.Println(k, v, string(v))
		}
		fmt.Println()
	}
}