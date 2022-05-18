package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("hello, welcome to cmd-calcApp")

	var firstNumber float64
	var secondNumber float64
	var operand string
	clearWindow()
	fmt.Print("firstNumber: ")
	fmt.Fscan(os.Stdin, &firstNumber)
	clearWindow()
	fmt.Print("operand: ")
	fmt.Fscanln(os.Stdin, &operand)
	clearWindow()
	fmt.Print("secondNumber: ")
	fmt.Fscan(os.Stdin, &secondNumber)
	clearWindow()
	calcApp(firstNumber, secondNumber, operand)
}

func calcApp(firstNumber float64, secondNumber float64, operand string) {
	var sum float64 = 0

	if operand == "+" {
		sum = firstNumber + secondNumber
	} else if operand == "-" {
		sum = firstNumber - secondNumber
	} else if operand == "*" {
		sum = firstNumber * secondNumber
	} else if operand == "/" {
		sum = firstNumber / secondNumber
	} else {
		fmt.Println("operand isn't correct: ", operand)
		operand = os.DevNull
	}
	if operand != os.DevNull {
		fmt.Println("Summary: ", firstNumber, operand, secondNumber, " = ", sum)
	}
}

func clearWindow() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}
