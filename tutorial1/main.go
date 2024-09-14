package main

import (
	"errors"
	"fmt"
)

func main() {
	// const somt string = "Hello World"
	// printme(somt)
	var division, remainder, err = divisor(12, 0)
	if err != nil {
		fmt.Print(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The division of integer i s %v", division)
	} else {
		fmt.Printf("The result of the division is %v and the remainder is %v", division, remainder)
	}
}

// func printme(input_str string) {
// 	fmt.Print(input_str)
// }

func divisor(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	return numerator / denominator, numerator % denominator, err
}
