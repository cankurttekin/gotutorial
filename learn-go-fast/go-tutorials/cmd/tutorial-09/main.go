package main
import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 3

func main() {
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"website1.com", "website2.com", "website3.com"}
	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)
}

func checkTofuPrices(website string, c chan string) {
	for {
		time.Sleep(1 * time.Second)
		var tofuPrice = rand.Float32() * 20
		if tofuPrice<= MAX_TOFU_PRICE {
			c <- website
			break
		}
	}
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(1 * time.Second)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	select {
	case website := <-chickenChannel:
		fmt.Printf("\nnotification sent: found a deal on chicken at: %s\n", website)
	case website := <-tofuChannel:
		fmt.Printf("\nemail sent: found a deal on tofu at: %s\n", website)
	}
}


