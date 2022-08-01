package main

import (
	"fmt"
	"sort"
)

type car struct {
	nTires int
	luxury bool
	nSeats int
	make   string
	model  string
	year   int
}

func main() {
	var i1 = 4
	var f1 = 2.0

	var result = i1 / int(f1)
	fmt.Println(result)
	fmt.Println(&i1)

	var myStrings [3]string
	myStrings[0] = "Cat"
	myStrings[1] = "Dog"
	myStrings[2] = "Fish"

	fmt.Println("First is:", myStrings[0])
	fmt.Printf("%v, %T\n", myStrings, myStrings)

	var myCar car
	myCar.make = "Honda"
	myCar.model = "CRV"
	myCar.year = 1997
	myCar.luxury = false
	myCar.nTires = 4
	myCar.nSeats = 5

	fmt.Printf("%v, %T\n", myCar, myCar)

	var caddilac = car{
		year:   1997,
		nSeats: 5,
		luxury: true,
	}
	fmt.Printf("%v, %T\n", caddilac, caddilac)

	var cars = []car{
		{make: "Honda", year: 2021},
		{make: "Toyota", year: 1997},
	}

	cars = append(cars, car{make: "Volvo", year: 2019})

	fmt.Printf("all: %v, %T\n", cars, cars)
	fmt.Println(len(cars))

	for e, i := range cars {
		fmt.Printf("%d: %v\n", e, i)
	}

	fmt.Println("Sorted?", sort.SliceIsSorted(cars, func(i, j int) bool {
		return cars[i].year < cars[j].year
	}))

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].year < cars[j].year
	})

	for e, i := range cars {
		fmt.Printf("%d: %v\n", e, i)
	}

	fmt.Println(" ==== M A P S ====")

	var m1 = make(map[string]int)
	m1["one"] = 1
	m1["two"] = 2
	m1["three"] = 3
	m1["four"] = 4
	m1["five"] = 5
	// fmt.Printf("%v, %T", m1, m1)
	for k, v := range m1 {
		fmt.Printf("%s: %d\n", k, v)
	}

	delete(m1, "four")
	_, ok := m1["four"]
	if ok {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
