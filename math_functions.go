package main

import "fmt"

func fibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	}

	fib := []int{0, 1}
	for i := 2; i < n; i++ {
		fib = append(fib, fib[i-1]+fib[i-2])
	}
	return fib
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	result := 1
	for i := 1; i < n; i++ {
		result *= i
	}
	return result
}

func is_prime(num int) bool {
	if num < 2 {
		return false
	}

	for i := 2; i < num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func print_hello(name string) {
	fmt.Println("Hello, " + name + "!")
}

func main() {
	fib := fibonacci(10)
	fmt.Println(fib)

	fact := factorial(5)
	fmt.Println(fact)

	prime := is_prime(17)
	fmt.Println(prime)

	print_hello("Alice")
}
