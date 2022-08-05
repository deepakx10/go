package main

import "fmt"

/*
func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Println("FizzBuzz")
			} else {
				fmt.Println("Fizz")
			}
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
*/
func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 3 && i%5 ==0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("buzz")
		}
		fmt.Println(i)
	}
}