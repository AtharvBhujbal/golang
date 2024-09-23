package main

import (
	"fmt"
	"slices"
	"strings"
)

// func main() {
// 	// const somt string = "Hello World"
// 	// printme(somt)
// 	var division, remainder, err = divisor(12, 0)
// 	if err != nil {
// 		fmt.Print(err.Error())
// 	} else if remainder == 0 {
// 		fmt.Printf("The division of integer i s %v", division)
// 	} else {
// 		fmt.Printf("The result of the division is %v and the remainder is %v", division, remainder)
// 	}
// }

// // func printme(input_str string) {
// // 	fmt.Print(input_str)
// // }

// func divisor(numerator int, denominator int) (int, int, error) {
// 	var err error
// 	if denominator == 0 {
// 		err = errors.New("cannot divide by zero")
// 		return 0, 0, err
// 	}
// 	return numerator / denominator, numerator % denominator, err
// }

func main() {
	// 1. Array
	//  a)Fixed length b)Same type c)Indexable d) contiguous in memory
	// var intArr [3]int32 = [3]int32{1, 2, 3}
	// intArr := [3]int32{1, 2, 4}
	intArr := [...]int32{1, 2, 3}
	fmt.Println(intArr)

	// 2. Slices
	// Slices are the vectors of golang
	intSlice := []int32{6, 7, 8}
	fmt.Println(intSlice)
	var intSlice2 []int32 = []int32{4, 5, 6}
	//can also use make function
	// var intSlice []int32 = make([]int32,3,10) //(datatype,length,capacity)
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)
	slices.Sort(intSlice)
	fmt.Println(intSlice)

	// 3. Maps
	// {'key':'value'}
	var myMap map[string]int32 = make(map[string]int32, 2)
	myMap2 := map[string]int32{"John": 32, "Atharv": 21, "Ayush": 18}
	fmt.Println(myMap, myMap2)
	var age, ok = myMap2["Atharv"]
	if !ok {
		fmt.Println("Invalide name!")
	} else {
		fmt.Printf("The age is %v \n", age)
	}

	for name, age := range myMap2 {
		fmt.Printf("Name: %v Age: %v \n", name, age) // Order is not preserved
	}

	//4. Strings and Runes
	//Strings are non-mutable
	//catStr[0] = 'a' cannot do this
	//Strings while indexings are the underlying representation of UTF-8 encoding this creats a indexing problem
	//It can be solved by Runes

	// var myString = "resume"
	var myString = []rune("resume")
	var indexed = myString[1]
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}

	//Inefficient way
	// var strSlice = []string{"s", "u", "b", "s", "c", "r", "i", "b", "e"}
	// var catStr string = ""
	// for i := range strSlice {
	// 	catStr += strSlice[i]
	// }
	// fmt.Printf("\n%v", catStr)

	//efficient way
	var strSlice2 = []string{"s", "u", "b", "s", "c", "r", "i", "b", "e"}
	var strBuilder strings.Builder
	for i := range strSlice2 {
		strBuilder.WriteString(strSlice2[i])
	}
	var catStr2 = strBuilder.String()
	fmt.Printf("\n%v", catStr2)

	//5. Structs and Interfaces
	// Struct - Defined as a data type with properties
	// Is a class of a golang.
	// Interfaces are the way to generalize the input parameter to take multiple data types
	// Is a method overloading in golang.

	type owner struct {
		name string
		age  int
	}

	type gasEngine struct {
		mpg     int
		gallons int
		owner
	}
	type engine interface {
		milesLeft() int
	}

	// func (e gasEngine) milesLeft() int{
	// 	return e.gallons * e.mpg
	// }

	// func canMakeIt(e gasEngine, miles int){
	// 	if miles<= e.milesLeft(){
	// 		fmt.Println("You Can make it there!!")
	// 	}else{
	// 		fmt.Println("Need to refuel!")
	// 	}
	// }

	// Above 'canMakeIt' method can be written as

	// func canMakeIt(e engine, miles int){ // gasEngine -> engine //accepting a interface type i.e. a structure having a milesLeft() property is compatible
	// 	if miles<= e.milesLeft(){
	// 		fmt.Println("You Can make it there!!")
	// 	}else{
	// 		fmt.Println("Need to refuel!")
	// 	}
	// }
	var myEngine gasEngine = gasEngine{25, 35, owner{"ALex", 5}}
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.name)

	//6. Pointers
	// Special types
	// Stores the value of memory location

	var p *int32 = new(int32)
	*p = 10     //Storing a value 10 at a location where the pointer is pointing to
	var i int32 //Regular var
	p = &i      //pointer is pointing to i
	*p = 2      //val of i also changes

	// Slices contain pointers to original data
	// Changing the val of copy changes the val of original slices
	var slice = []int32{1, 2, 3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice) // Should be [1,2,3] but it is [1,2,4]
	fmt.Println(sliceCopy)
}
