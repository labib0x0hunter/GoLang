package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getParseError(err1, err2 error) error {
	if err1 != nil && err2 != nil {
		return fmt.Errorf("%v and %v", err1, err2)
	} else if err1 != nil {
		return fmt.Errorf("%v", err1)
	} else if err2 != nil{
		return fmt.Errorf("%v", err2)
	}
	return nil
}

func validateInput(input string) (a, b float64, operation string, err error) {
	input = strings.TrimSpace(input)
	splitInput := strings.Split(input, " ")
	if len(splitInput) != 3 {
		err = fmt.Errorf("invalid operation")
		return
	}
	a, parseErrorA := strconv.ParseFloat(splitInput[0], 64)
	b, parseErrorB := strconv.ParseFloat(splitInput[2], 64)
	operation = splitInput[1]
	err = getParseError(parseErrorA, parseErrorB)
	return
}

func add(a, b float64) {
	printResult(a, b, a+b, "+")
}

func subtract(a, b float64) {
	printResult(a, b, a-b, "-")
}

func multiply(a, b float64) {
	printResult(a, b, a*b, "*")
}

func division(a, b float64) {
	if b == 0 {
		printError(fmt.Errorf("can't divide by zero"))
		return
	}
	printResult(a, b, a/b, "/")
}

func printError(err error) {
	fmt.Println("Error:", err)
}

func printResult(a, b, result float64, operation string) {
	fmt.Println(a, operation, b, "=", result)
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Printf("Give operation: ")
		input, err := reader.ReadString('\n')

		if err != nil {
			printError(err)
			continue
		}

		a, b, operation, err := validateInput(input)
		if err != nil {
			printError(err)
			continue
		}

		switch operation {
		case "+":
			add(a, b)
		case "-":
			subtract(a, b)
		case "*":
			multiply(a, b)
		case "/":
			division(a, b)
		default:
			printError(fmt.Errorf("unknown operation"))
		}
	}
}
