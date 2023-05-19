package main

import (
	"github.com/inancgumus/screen"
	"calculator/operations"	
	"fmt"
)

func main() {
	displayMenu()
}

func displayMenu() {
	option := 1 

	for option!=0 {
		fmt.Println("\n\nChoose a number for an operation:")
		fmt.Println("[1] - Addition")
		fmt.Println("[2] - Subtraction")
		fmt.Println("[3] - Multiplication")
		fmt.Println("[4] - Division")
		fmt.Println("[0] - Exit")
		
		fmt.Scanln(&option)
		if option == 0 {
			return
		}
		startOperation(option)
	}
	
}

func startOperation(option int) {
	if option == 1 {
		num1, num2 := getNumbers()
		result := operations.Add(num1, num2)
		screen.Clear()
		fmt.Printf("%d + %d = %d", num1, num2, result)
	} else if option == 2 {
		num1, num2 := getNumbers()
		result := operations.Subtract(num1, num2)
		screen.Clear()
		fmt.Printf("%d - %d = %d", num1, num2, result)
	} else if option == 3 {
		num1, num2 := getNumbers()
		result := operations.Multiply(num1, num2)
		screen.Clear()
		fmt.Printf("%d * %d = %d", num1, num2, result)
	} else if option == 4 {
		num1, num2 := getNumbers()
		result := operations.Divide(num1, num2)
		screen.Clear()
		fmt.Printf("%d / %d = %.2f", num1, num2, result)
	} else {
		fmt.Print("Incorrect option. Select an option from the menu.")
	} 
}

func getNumbers() (int, int) {
	var num1, num2 int

	fmt.Println("Input first integer: ")
	fmt.Scanln(&num1)
	fmt.Println("Input second integer: ")
	fmt.Scanln(&num2)

	return num1, num2
}