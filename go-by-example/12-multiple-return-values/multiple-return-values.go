// built in support for multiple return values
// for ex. return both result and err

package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a, b)

	_, c := vals()
	fmt.Println(c)
}

