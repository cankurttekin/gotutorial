// associative data type
// like hash or dicts
// make(map[key-type]val-type)
// if key not exist, it returns zero value of value type
// len, delete, clear

package main

import (
	"fmt"
	"maps"
)

func main() {
	m := map make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

