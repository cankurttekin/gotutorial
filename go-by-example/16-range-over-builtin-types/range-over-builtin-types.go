// range iterates over elements in a variety of built-in data structures 

package main

import "fmt"

func main() {
	// range to sum numbers in a slice

	// range on arrays and slices provides both the index and value for each entry
	// we dont need index below so we ignore it with the blank identifier _
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)


	// when we need index
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// range on maps iterates over key/value pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}


	// range can also iterate over just the keys or values of a map
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range on strings iterates over Unicode code points (runes)
	for i, c := range "go" {
		fmt.Println(i, c) // i is the index, c is the rune value
	}
}

