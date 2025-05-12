package main

import "fmt"

func main() {
	var intArr [5]int32
	fmt.Println(intArr[0])

	fmt.Println(&intArr[0])

	var initArr [5]int32 = [5]int32{1, 2, 3, 4, 5}
	fmt.Println(initArr)

	initArr2 := [2]int32{1, 2}
	fmt.Println(initArr2)
	
	var intSlice []int32 = []int32{4,5,6}
	fmt.Printf("the length of the slice is %v, capacity is %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("the length of the slice is %v, capacity is %v\n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)

}
