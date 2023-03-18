package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var userInput int
	var secretNumber int
	var quit bool = false

	rand.Seed(time.Now().Unix())
	secretNumber = rand.Intn(10)
	fmt.Println("Secret number is:", secretNumber)

	fmt.Println("Please enter a number:")
	fmt.Scan(&userInput)
	fmt.Println("You entered:", userInput)

	for quit != true {
		if userInput == secretNumber {
			fmt.Println("You won!")
		} else if userInput < secretNumber {
			fmt.Println("Too low...")
		} else if userInput > secretNumber {
			fmt.Println("Too High...")
		}
	}
}
