// Go will infer the type of initialized variables.
// Variables declared without a corresponding initialization are zero-valued.
// The := syntax is shorthand for declaring and initializing a variable. This syntax is only available inside functions.

package main

import "fmt"

func main() {
	var a = "my initial value"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = false 
	fmt.Println(d)

	var e int
	fmt.Println(e) // will print 0

	x := 42
	fmt.Println(x)
}

