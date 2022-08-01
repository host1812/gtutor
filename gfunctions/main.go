package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Alive bool
}

func main() {
	fmt.Println(addTwoInts(3, 5))
	fmt.Println(addMany(3, 5, 2))
	p1 := Person{
		Name:  "Alex",
		Age:   97,
		Alive: true,
	}

	p2 := p1.aged(1)
	fmt.Println(p1)
	fmt.Println(p2)

}

func (p Person) aged(years int) Person {
	p.Age += 1
	return p
}

func addTwoInts(a, b int) int {
	return a + b
}

func addMany(nums ...int) int {
	total := 0
	for _, x := range nums {
		total = total + x
	}
	return total
}
