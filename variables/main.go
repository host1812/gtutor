package main

import (
	"bufio"
	"fmt"
	"os"
)

const prompt = "and press ENTER"

func main() {
	var theNumber int
	theNumber = 1
	var anotherNumber = 2
	superNumber := 42

	// var answer int
	fmt.Println("Guess the Number")
	fmt.Println("----------------")
	fmt.Println("Enter a number ", prompt)
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	fmt.Println("Multiply by ", theNumber, prompt)
	reader.ReadString('\n')
	fmt.Println("Multiply by ", anotherNumber, prompt)
	reader.ReadString('\n')
	fmt.Println("Multiply by ", superNumber, prompt)
}
