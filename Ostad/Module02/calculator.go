package main

import (
	"fmt"
	"os"
)

func main() {

	var (
		a, b, result float64
		operation    string
	)

	fmt.Printf("Give two number : ")
	_, errInput := fmt.Scanf("%f %f", &a, &b)

	if errInput != nil {
		fmt.Println("Error: ", errInput)
		os.Exit(0)
	}

	fmt.Printf("Operation: ")
	_, errOperation := fmt.Scanf("%s", &operation)

	if errOperation != nil {
		fmt.Println("Error: ", errOperation)
		os.Exit(0)
	}

	switch operation {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			fmt.Println("Can't Divide by zero")
			os.Exit(0)
		}
		result = a / b
	default:
		fmt.Println("Unknown Operation")
	}

	fmt.Println(a, operation, b, "=", result)
}
