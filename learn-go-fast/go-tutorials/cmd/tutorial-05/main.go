package main
import (
	"fmt"
	"strings"
)

func main() {
	var myString = "cankurttekin"
	var indexed = myString[0]
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v := range myString {
		fmt.Printf("index: %d, value: %c\n", i, v)
	}
	
	var myRune = 'c' 
	var myRune2 = 'ç'
	var myRune3 = []rune{'c', 'ç', 'ğ', 'ü', 'ş', 'ı'}
	var myString2 = []rune(myString)
	fmt.Printf("myRune: %c, %T\n", myRune, myRune)
	fmt.Printf("myRune2: %c, %T\n", myRune2, myRune2)
	fmt.Printf("the length of myRune3 is %v", len(myRune3))
	fmt.Printf("the length of myString2 is %v", len(myString2))

	var strSlice = []string{"a", "b", "c"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v", catStr)
}

