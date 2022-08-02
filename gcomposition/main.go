package main

import (
	"fmt"

	"github.com/host1812/gtutor/gcomposition/vehicles"
)

func main() {
	fmt.Println("started")
	c1 := vehicles.Car{
		Make:  "Volvo",
		Model: "XC90",
		Year:  2017,
		Vehicle: vehicles.Vehicle{
			NWheels: 3,
			NSeats:  19,
		},
	}
	c1.ShowDetails()
	fmt.Println("finished")
}
