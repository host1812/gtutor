package vehicles

import "fmt"

type Vehicle struct {
	NWheels int
	NSeats  int
}

func (v *Vehicle) ShowDetails() {
	fmt.Printf(
		"Seats: %d, Wheels: %d\n",
		v.NSeats,
		v.NWheels,
	)
}
