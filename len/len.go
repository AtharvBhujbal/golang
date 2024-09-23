package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var str = "résumé"

	var length = utf8.RuneCountInString(str)

	fmt.Println("Length of the string is :", length)   // Rune calculates the length of byte array by merging two val if they are representing a special character(Ex: é uses two bytes to represent itself [110101,010110])
	fmt.Println("Length of the string is :", len(str)) // Len function calculates the length of byte array of utf8 encoding [,.,.,.,.,]
}
