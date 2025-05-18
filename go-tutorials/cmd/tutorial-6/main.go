package main
import "fmt"

type gasEngine struct {
	mpg uint8
	gallons uint8
	owner owner
	int
}

type electricEngine struct {
	mpkwh uint8
	kwh uint8
}

type owner struct {
	name string
	license string
}

func (g gasEngine) totalMiles() uint8 {
	return g.mpg * g.gallons
}

func (ev electricEngine) totalMiles() uint8 {
	return ev.mpkwh * ev.kwh
}


type engine interface {
	totalMiles() uint8

func main() {
	var myEngine gasEngine = gasEngine{mpg: 25, gallons: 10, owner: owner{name: "John Doe", license: "12345"}, int: 1}
	var myElectricEngine electricEngine = electricEngine{mpkwh: 4, kwh: 10}

	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.owner)
	fmt.Println(myElectricEngine.mpkwh, myElectricEngine.kwh)


	// structs also has methods
	fmt.Printf("total miles left in tank: %v", myEngine.totalMiles())
	fmt.Printf("\ntotal miles left in battery: %v", myElectricEngine.totalMiles())


}
