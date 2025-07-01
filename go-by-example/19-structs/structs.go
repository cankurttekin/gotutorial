// structs are typed collections of fields
// go is a garbage collected language
// so you can safely return a pointer to a local var
// when there are no active references to it it will be cleaned up


package main

import "fmt"

type person struct {
	name string
	age int
}

func newPerson(name string) *person {
	p := person{ name: name }
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	

	//fmt.Println(person{"Fred"})
	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s)
	fmt.Println(s.name)

	// pointers are automatically dereferenced
	sp := &s
	fmt.Println(sp)
	fmt.Println(sp.name)

	// structs are mutable
	// if a struct type is only used for a single value
	// we dont have to give it a name
	// the value can have an anonymous struct type
	dog := struct {
		name string
		isGoodDog bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)



}










