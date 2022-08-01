package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

var keyPressedChan chan rune

func main() {
	// go doS("Hellow World")
	// fmt.Println("goroutine started")
	// for {

	// }
	keyPressedChan = make(chan rune)

	go listenKeyPressed()

	fmt.Println("Press and key, or q to quit.")
	_ = keyboard.Open()
	defer func() {
		keyboard.Close()
	}()
	for {
		char, _, _ := keyboard.GetSingleKey()
		if char == 'q' || char == 'Q' {
			break
		}
		keyPressedChan <- char
	}
}

func listenKeyPressed() {
	for {
		key := <-keyPressedChan
		fmt.Println("You pressed:", string(key))
	}
}

func doS(s string) {
	nRun := 0
	for {
		fmt.Printf("s: %s (%d)\n", s, nRun)
		nRun += 1
		if nRun > 4 {
			break
		}
	}
}
