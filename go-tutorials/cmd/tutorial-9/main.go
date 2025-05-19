package main
import (
	"fmt"
	"time"
)

func main() {
	var c = make(chan int, 5) // create a buffered channel with a capacity of 5
	go process(c) // start a goroutine
	for i := range c {
		fmt.Println(i)
		time.Sleep(1 * time.Second) // simulate some work
	}
}

func process(c chan int) {
	defer close(c) // ensure the channel is closed when done
	for i := 0; i < 5; i++ {
		c <- i
	}
	fmt.Println("exiting process")
}
