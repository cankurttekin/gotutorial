package main
import "fmt"
import "unicode/utf8"

func main() {
	var intVar int = 42
	fmt.Println(intVar)

	var floatVar float32 = 3.14
	var intVar32 int32 = 2
	var result float32 = floatVar + float32(intVar32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)

	var myString string = "Hello, Go!"
	fmt.Println(myString)
	fmt.Println("Length of string:", len(myString))
	fmt.Println("Length of string:", utf8.RuneCountInString(myString))

	var myBool bool = true
	fmt.Println(myBool)

	var intVar1 int
	fmt.Println(intVar1)

	myAnotherString := "Hello, Go!"
	fmt.Println(myAnotherString)

	var var1, var2, var3 int = 1, 2, 3
	fmt.Println(var1, var2, var3)

	const myConst int = 100
	// myConst = 200 // Uncommenting this line will cause a compilation error
	fmt.Println(myConst)
}
