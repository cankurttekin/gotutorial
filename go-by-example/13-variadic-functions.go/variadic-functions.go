// can be called with any num of trailing arguments
// for example fmt.Println

package main

import "fmt"

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}


func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1,2,3,4}
	sum(nums...)

	sum([]int{99, 100}...)
}


