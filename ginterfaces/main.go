package main

import "fmt"

type Animal interface {
	Says() string
	HowManyLegs() int
}

type Dog struct {
	Name  string
	Sound string
	NLegs int
}

func (d *Dog) Says() string {
	return d.Sound
}

func (d *Dog) HowManyLegs() int {
	return d.NLegs
}

type Cat struct {
	Name    string
	Sound   string
	NLegs   int
	HasTail bool
}

func (c *Cat) Says() string {
	return c.Sound
}

func (c *Cat) HowManyLegs() int {
	return c.NLegs
}

func main() {
	d1 := Dog{
		Name:  "pp",
		NLegs: 4,
		Sound: "wiwi",
	}
	Riddle(&d1)

	c1 := Cat{
		Name:  "mm",
		NLegs: 4,
		Sound: "mya",
	}
	Riddle(&c1)
}

func Riddle(a Animal) {
	riddle := fmt.Sprintf(
		`This animal says %s and has %d legs. What animal is it?`,
		a.Says(),
		a.HowManyLegs(),
	)
	fmt.Println(riddle)
}
