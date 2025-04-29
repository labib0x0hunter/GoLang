package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	for {

		// // fmt Input
		// var (
		// 	a, b           float64
		// 	operation         string
		// 	newLineOnBadInput string
		// )
		// fmt.Printf("Give two number : ")
		// input, errInput := fmt.Scanf("%f %f", &a, &b)

		// if errInput != nil || input != 2{
		// 	fmt.Scanln(&newLineOnBadInput)
		// 	fmt.Println("Error: ", errInput)
		// 	// os.Exit(0)
		// 	continue
		// }

		// fmt.Printf("Operation: ")
		// _, errOperation := fmt.Scanf("%s", &operation)

		// if errOperation != nil {
		// 	fmt.Println("Error: ", errOperation)
		// 	// os.Exit(0)
		// 	continue
		// }

		// bufio Reader input
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Give operation: ")
		input, inputError := reader.ReadString('\n')

		if inputError != nil {
			fmt.Println("Error: ", inputError)
			continue
		}

		input = strings.Trim(input, "\n")
		splitInput := strings.Split(input, " ")
		if len(splitInput) != 3 {
			fmt.Println("Error: Invalid Operation")
			continue
		}
		a, parseErrorA := strconv.ParseFloat(splitInput[0], 64)
		b, parseErrorB := strconv.ParseFloat(splitInput[2], 64)
		operation := splitInput[1]

		if parseErrorA != nil {
			fmt.Println("Error: ", parseErrorA)
			continue
		}
		if parseErrorB != nil {
			fmt.Println("Error: ", parseErrorB)
			continue
		}

		var result float64

		switch operation {
		case "+":
			result = a + b
		case "-":
			result = a - b
		case "*":
			result = a * b
		case "/":
			if b == 0 {
				fmt.Println("Error: Can't Divide by zero")
				// os.Exit(0)
				continue
			}
			result = a / b
		default:
			fmt.Println("Unknown Operation")
			continue
		}

		fmt.Println(a, operation, b, "=", result)
	}
}
