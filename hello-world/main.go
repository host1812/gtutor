package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	var simpleSay string
	simpleSay = "Testing test"

	anotherSay := "Another say"
	sayHello(simpleSay)
	sayHello(anotherSay)
}

func sayHello(what string) {
	fmt.Println(what)
}
