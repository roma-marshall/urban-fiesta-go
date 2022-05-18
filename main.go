package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("go run main.go hello world")

	var firstNumber float64
	var secondNumber float64
	var operand string
	var sum float64
	sum = 0

	fmt.Print("firstNumber: ")
	fmt.Fscan(os.Stdin, &firstNumber)

	fmt.Print("operand: ")
	fmt.Fscanln(os.Stdin, &operand)

	fmt.Print("secondNumber: ")
	fmt.Fscan(os.Stdin, &secondNumber)

	if operand == "+" {
		sum = firstNumber + secondNumber
	} else if operand == "-" {
		sum = firstNumber - secondNumber
	} else if operand == "*" {
		sum = firstNumber * secondNumber
	} else if operand == "/" {
		sum = firstNumber / secondNumber
	}

	fmt.Println("Summary: ", firstNumber, operand, secondNumber, " = ", sum)
}
