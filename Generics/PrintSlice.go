package main

import (
	"fmt"
)

// Print slice of any data types
func PRINT[T any](arr []T) {
	for _, a := range arr {
		fmt.Print(a, " ")
	}
	fmt.Println()
}

func main() {
	// arr is a slice of int data types 
	arr := []int{1, 2, 3}
	PRINT(a)

	// brr is a slice of string data types
	brr := []string{"a", "b", "c"}
	PRINT(b)
}
