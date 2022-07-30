package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	coffees := make(map[int]string)

	coffees[1] = "Some1"
	coffees[2] = "Some2"
	coffees[3] = "Some3"
	coffees[4] = "Some4"
	coffees[5] = "Some5"

	fmt.Println("Menu")
	fmt.Println("----")
	fmt.Println("1. Some1")
	fmt.Println("2. Some2")
	fmt.Println("3. Some3")
	fmt.Println("4. Some4")
	fmt.Println("5. Some5")
	fmt.Println("Q. Quit")

	fmt.Println("Press a key... Press ESC to quit...")
	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		i, _ := strconv.Atoi(string(char))
		t := fmt.Sprintf("You chose: %s", coffees[i])
		fmt.Println(t)
		if char == 'q' || char == 'Q' {
			break
		}
	}
	fmt.Println("Exiting.")
}
