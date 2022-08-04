package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 10; i++ {
		fmt.Print(i)
	}
	fmt.Println()

	y := 1
	for y < 10 {
		fmt.Print(y)
		y++
	}
	fmt.Println()

	// for j := 0; j < 10; j++ {
	// 	fmt.Print(j)
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Print(i)
	// 	}
	// 	fmt.Println()
	// }
	// reader := bufio.NewReader(os.Stdin)
	// ch := make(chan string)
	// go logger.Listen(ch)
	// fmt.Println("Enter something")
	// for i := 0; i < 5; i++ {
	// 	fmt.Print("> ")
	// 	input, _ := reader.ReadString('\n')
	// 	ch <- input
	// 	time.Sleep(time.Second)
	// }
}
