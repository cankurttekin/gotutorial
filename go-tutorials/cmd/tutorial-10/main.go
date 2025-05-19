package main

import "fmt"

type contactInfo struct {
	Name		string
	Email		string
}

type purchaseInfo struct {
	Name		string
	Price		float32
	amount		int
}

type gasEngine struct {
	gallons		float32
	mpg			float32
}

type electricEngine struct {
	kWh			float32
	mpkwh		float32
}

// generics with structs
type car [T gasEngine | electricEngine] struct {
	carMake		string
	carModel	string
	engine T
}

func main() {
	var intSlice = []int{1, 2, 3, 4, 5}
	fmt.Println(isEmpty[int](intSlice))
	fmt.Println(sumSlice[int](intSlice))

	var float32Slice = []float32{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println(isEmpty[float32](float32Slice))
	fmt.Println(sumSlice[float32](float32Slice))

	var emptySlice []int
	fmt.Println(isEmpty(emptySlice))

	//
	var contacts []contactInfo = loadJSON[contactInfo]("contacts.json")
	fmt.Printf("\n%+v\n", contacts)

	var purchases []purchaseInfo = loadJSON[purchaseInfo]("purchases.json")
	fmt.Printf("\n%+v\n", purchases)

	//
	var gasCar = car[gasEngine]{
		carMake: "Toyota",
		carModel: "Camry",
		engine: gasEngine{
			gallons: 10,
			mpg: 30,
		},
	}
	var electricCar = car[electricEngine]{
		carMake: "Tesla",
		carModel: "Model S",
		engine: electricEngine{
			kWh: 100,
			mpkwh: 4,
		},
	}
}

// generics with functions
func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

func loadJSON[T contactInfo | purchaseInfo](filePath string) []T {
	data, _ = ioutil.ReadFile(filePath)

	var loaded = []T{}
	json.Unmarshal(data, &loaded) // unmarhal means to convert the JSON data into a Go struct

	return loaded
}

