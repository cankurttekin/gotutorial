// go supports anon funcs
// anon funcs useful when you want to define a func inline without having to name it

package main

import "fmt"

// this func returns another func
// which we define anonymously in the body
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}


func main() {
	nextInt := intSeq()

	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	newInts := intSeq()
	fmt.Println(newInts()) // 1
}
