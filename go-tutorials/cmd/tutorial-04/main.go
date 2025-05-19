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
	fmt.Println(initArr2[0])
	
	var intSlice []int32 = []int32{4,5,6}
	fmt.Printf("the length of the slice is %v, capacity is %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("the length of the slice is %v, capacity is %v\n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)

	var intSlice2 = []int32{1, 2, 3}
	intSlice = append(intSlice2, intSlice...)
	fmt.Println(intSlice)


	// you can also specify the length and capacity of a slice
	var intSlice3 []int32 = make([]int32, 5, 10)
	fmt.Printf("the length of the slice is %v, capacity is %v\n", len(intSlice3), cap(intSlice3))


	// maps
	var myMap map[string]uint32 = make(map[string]uint32) // this means keys are string and values are uint32
	fmt.Println(myMap)

	var myMap2 = map[string]uint32{"John": 1, "Jane": 2}
	fmt.Println(myMap2)

	var age, ok = myMap2["Can"]
	delete(myMap2, "John")
	if ok {
		fmt.Println("the age of Can is", age)
	} else {
		fmt.Println("Can is not in the map")
	}

	for name := range myMap2 {
		fmt.Println(name)
	}

	for name, age := range myMap2 {
		fmt.Println(name, age)
	}
	
	for i, v := range initArr {
		fmt.Printf("index: %v, val: %v \n", i, v)
	}


	// while loop
	var i int32 = 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	
}
