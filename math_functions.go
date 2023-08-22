package main

import "fmt"

// I check for:
// Readability
// Understandability
// Performance
// Error handling
// Bugs & Logical flaws
// DRY-Principle -> Don't write redundant code
// SOP-Principle -> Each thing does its own thing
// KISS-Principle -> Try to keep it simple
// SOLID-Principle -> https://en.wikipedia.org/wiki/SOLID
// POLA-Principle -> Code should work as one would expect
// YAGNI-Principle -> Dont code stuff you might not need

// Removed abbreviations to make code easier to read and understand
// Removed append and used a more simplified approach which should be less memory intensive
// Code wouldn't work if length was value 1. Added a check to catch that case.
func fibonacci(length int) []int {
	if length <= 0 {
		return []int{}
	}

	if length == 1 {
		return []int{0}
	}

	fibonacciSequence := make([]int, length)
	fibonacciSequence[0], fibonacciSequence[1] = 0, 1
	for index := 2; index < length; index++ {
		fibonacciSequence[index] = fibonacciSequence[index-1] + fibonacciSequence[index-2]
	}

	return fibonacciSequence
}

// Removed abbreviations to make code easier to read and understand
// for-loop was flawed which lead to an incorrect factorial. Fixed it by replacing = with <=
// Added error handling for negative numbers
func factorial(number int) (int, error) {
	if number < 0 {
		return 0, fmt.Errorf("invalid input: can't calculate factorial with negative numbers")
	}

	if number == 0 {
		return 1, nil
	}

	factorialResult := 1
	for index := 1; index <= number; index++ {
		factorialResult *= index
	}

	return factorialResult, nil
}

// Renamed function according to best practices https://go.dev/doc/effective_go#mixed-caps
// Removed abbreviations to make code easier to read and understand
func isPrime(number int) bool {
	if number < 2 {
		return false
	}

	for divisor := 2; divisor < number/2; divisor++ {
		if number%divisor == 0 {
			return false
		}
	}

	return true
}

// Renamed function according to best practices https://go.dev/doc/effective_go#mixed-caps
// Instead of implementing error handling, an empty string will instead produce a hello message without personalization
// Used sprintf instead of + string concatenation, although sprintf may be less performant.
func printHello(name string) {
	var message string
	if len(name) > 0 {
		message = fmt.Sprintf("Hello, %s!", name)
	} else {
		message = "Hello!"
	}

	println(message)
}

// No need to declare variables if the only use case is to return a function value.
func main() {
	fmt.Println(fibonacci(10))

	fmt.Println(factorial(5))

	fmt.Println(isPrime(43))

	printHello("Alice")
}
