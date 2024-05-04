package main

import (
	"fmt"
)

// Function to calculate factorial
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var num int

	fmt.Println("Factorial Calculator")
	fmt.Print("Enter a non-negative integer: ")
	fmt.Scanln(&num)

	if num < 0 {
		fmt.Println("Error: Please enter a non-negative integer.")
		return
	}

	result := factorial(num)
	fmt.Printf("Factorial of %d is %d\n", num, result)
}
