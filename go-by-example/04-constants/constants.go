// a numeric constant has no type until it is given one

package main

import (
	"fmt"
	"math"
)

const s string = "my constant string"

func main() {
	fmt.Println(s)

	const x = 500000000

	const d = 3e20 / x
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(x))
}
