package main

import "fmt"

func main() {
	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
	for _, word := range words {
		//fmt.Println(word)
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word")
		case 5:
			fmt.Println(word, "is right sized", size)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word")
		}
	}
}