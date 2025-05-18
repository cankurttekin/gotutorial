package main

import (
	"fmt"
)

func main() {
	// pointers
	var p *int32 // will hold mem addr of an int32 pointer initially nil
	p = new(int32) // p now points to a new int32 variable
	var i int32

	fmt.Printf("\nthe val p points to is %v", *p) // dereference p to get the value it points to
	fmt.Printf("\nthe value of i is %v", i) // i is 0 by default

	*p = 10 // set the value of the int32 variable that p points to
	fmt.Printf("\nthe val p points to is %v", *p) // dereference p to get the value it points to
	fmt.Printf("\nthe value of i is %v", i) // i is still 0
	var k int32 = 2
	i = k
	fmt.Printf("\nthe value of i is %v", i) // i is now 2


	// funcs with pointers
	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nthe mem loc of thing1 arr is %p", &thing1)
	var result [5]float64 = square(thing1)
	fmt.Printf("\nthe result is %v", result)


	var result2 [5]float64 = square2(&thing1)
	fmt.Printf("\nthe result2 is %v", result2)
	fmt.Printf("\nthe val of thing1 arr is %v", thing1)
	fmt.Printf("\nthe mem loc of result2 arr is %p", &result2)
}


func square(thing2 [5]float64) [5]float64 {
	fmt.Printf("\nthe mem loc of thing2 arr is %p", &thing2) 
	// mem locs are different this means
	// these are different arrays
	// we can modify the vals of thing2 without affecting thing1
	// but we are also doubling the memory usage
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return thing2
}


// we can also pass a pointer to the array to the function
func square2(thing3 *[5]float64) [5]float64 {
	fmt.Printf("\nthe mem loc of thing3 arr is %p", thing3) 
	// mem locs are the same this means
	// these are the same arrays
	// we can modify the vals of thing3 and it will affect thing1
	// but we are not doubling the memory usage
	for i := range *thing3 {
		thing3[i] = thing3[i] * thing3[i]
	}
	return *thing3
}









