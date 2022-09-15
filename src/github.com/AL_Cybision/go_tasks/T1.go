// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Car struct {
	tyres  int
	model  string
	colour string
	price  int
}

func car_details_print(mycar Car) {

	fmt.Println(mycar.tyres)
	fmt.Println(mycar.model)
	fmt.Println(mycar.colour)
	fmt.Println(mycar.price)

}

func main() {

	var car1 Car

	car1.tyres = 8
	car1.model = "Jeep - XRO"
	car1.colour = "Jet - Black"
	car1.price = 900000

	car_details_print(car1)

}
