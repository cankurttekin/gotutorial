// slices are typed only by the elements they contain, not by their size
// an uninitialized slice is nil and len is 0
// to create with non-zero length and capacity use make

package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "learning"
	s[1] = "go"
	s[2] = "slices"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))


	s = append(s, "is", "fun")
	s = append(s, "i", "guess")
	fmt.Println("app:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"a", "b", "c", "d", "e"}
	fmt.Println("dcl:", t)

	t2 := []string{"a", "b", "c", "d", "e"}
	if slices.Equal(t, t2) {
		fmt.Println("slices are equal")
	} else {
		fmt.Println("slices are not equal")
	}

	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}

