// fact func calls itself until it reaches the base case fact(0)
// anon functions can be also recursive
// but this requires explicitly declaring a variable
// with var to store the function before its defined 
package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println("Factorial 7:", fact(7))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println("Fibonacci 7:", fib(7))
}

