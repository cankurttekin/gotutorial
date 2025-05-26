package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("write, ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon.")
	default:
		fmt.Println("It's after noon.")
	}


	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool and my value is", t)
		case int:
			fmt.Println("I'm an int and my value is", t)
		default:
			fmt.Println("I don't know what I am")
		}
	}
	whatAmI(true)
	whatAmI(42)
	whatAmI("hello")
}







