package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type user struct {
	name   string
	age    int
	favnum float64
}

func main() {
	reader = bufio.NewReader(os.Stdin)
	var user user
	user.name = readString("Name?")
	user.age = readInt("Age?")
	user.favnum = readFloat("Fav Num?")
	fmt.Printf("User's name: %s, age: %d, favnum: %.2f.\n",
		user.name,
		user.age,
		user.favnum,
	)
}

func prompt() {
	fmt.Print("> ")
}

func readString(s string) string {
	fmt.Println(s)
	prompt()
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)
	return userInput
}

func readInt(s string) int {
	fmt.Println(s)
	prompt()
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)
	num, err := strconv.Atoi(userInput)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func readFloat(s string) float64 {
	fmt.Println(s)
	prompt()
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)
	num, err := strconv.ParseFloat(userInput, 64)
	if err != nil {
		log.Fatal(err)
	}
	return num
}
