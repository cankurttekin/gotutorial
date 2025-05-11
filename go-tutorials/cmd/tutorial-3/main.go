package main
import (
	"errors"
	"fmt"
)

func main() {
	var printVal = "Hello, World!"
	printMe(printVal)

	var num1 int = 11
	var den1 int = 2
	var result, remainder, err int = intDivision(num1, den1)
	if err != nil {
		fmt.Println("Error:", err)
	} else if result == 0 {
		fmt.Println("Result is zero")
	} else {
		fmt.Printf("Result: %v, Remainder: %v\n", result, remainder)
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("division by zero")
		return 0, 0, err
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}

