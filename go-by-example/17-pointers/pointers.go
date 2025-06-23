package main

import "fmt"

// arguments passed to this func by value
func zeroval(ival int) {
	ival = 0
}

// *int means that the argument is a pointer to an int
// *iptr in func body dereferences the pointer from its mem add. to
// the current val at that add
// assigning a value to a deferenced pointer(*iptr) changes the value at the referenced add.
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i) // pass the address of i
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i) // prints the memory address of i
}
