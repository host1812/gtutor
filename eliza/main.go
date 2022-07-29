package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/host1812/gtutor/eliza/doctor"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(doctor.Intro())
	for {
		fmt.Print("> ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)
		if userInput == "quit" {
			break
		}
		response := doctor.Response(userInput)
		fmt.Println(response)
	}
}
