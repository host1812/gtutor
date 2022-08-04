package main

import (
	"fmt"
	"math/rand"
	"time"
)

var ch1 = make(chan string)
var ch2 = make(chan string)

func main() {

	go task1()
	go task2()

	// for {
	// 	play()
	// 	time.Sleep(time.Second)
	// }
	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}

func task1() {
	for {
		time.Sleep(time.Second)
		ch1 <- "one"
	}
}

func task2() {
	for {
		time.Sleep(time.Second * 2)
		ch2 <- "two"
	}
}

func play() {
	rand.Seed(time.Now().UnixNano())

	switch value := rand.Intn(10); {
	case value < 9:
		fmt.Println("< 9")
	case value < 10:
		fmt.Println("< 10")
	default:
		fmt.Println("try again")
	}
}
