package main

import "fmt"

// Generics are "method overloading" in go
// You can specify the type of input data type in a square brackets
// New data type "any" can also be used
// It can not be used for the sumSlice function as everything is not compatible with addition operator
// Ex : we cannot use addition operator with string (strings are immutable in go)

func main() {
	var intSlice = []int32{5, 6, 7}
	fmt.Println(sumSlice(intSlice))
	var floatSlice1 = []float32{3.8, 4.7, 2.1}
	fmt.Println(sumSlice(floatSlice1))
	var floatSlice2 = []float64{9.1611, 7.2723, 9.133773}
	fmt.Println(sumSlice(floatSlice2))
	fmt.Println(isEmpty(floatSlice1))

}
func sumSlice[T int32 | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}
