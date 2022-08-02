package vehicles

import "fmt"

type Car struct {
	Make       string
	Model      string
	Year       int
	IsElectric bool
	IsHybrid   bool
	Vehicle    Vehicle
}

func (c *Car) ShowDetails() {
	fmt.Printf(
		"Make: %s\n",
		c.Make,
	)
	c.Vehicle.ShowDetails()
}
