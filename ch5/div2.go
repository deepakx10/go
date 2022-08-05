package main

import (
	"errors"
	"fmt"
	"os"
)

func divAndRem (numerator int, denominator int) (result int, remainder int, err error) {
	if denominator == 0 {
		err = errors.New("Can't devide by zero")
		return result, remainder, err
	}	
	result, remainder = numerator/denominator, numerator%denominator
	return result, remainder, err
}

func main() {
	//result, reminder, err := divAndRem(5, 0) 
		//fmt.Println(result, reminder, err)
	result, remainder, err := divAndRem(8, 4)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)
}