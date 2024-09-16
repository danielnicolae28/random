package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
)

func enterNum() int {
	var randomNum string
	fmt.Print("Enter a number: ")
	fmt.Scan(&randomNum)
	num, err := strconv.Atoi(randomNum)
	if err != nil {
		fmt.Println("You did not enter a numer, please enter a number")
		num = enterNum()
	}
	return num
}

func randomNumber() {
	var i int = 0
	for {

		num := enterNum()
		randNum := rand.IntN(10)
		if num > 10 {
			fmt.Println("Number is to big, please enter a number < 10")
			num = enterNum()
		}

		fmt.Printf("your number : %v, random: %v\n", num, randNum)
		if num != randNum {
			i++
			continue
		} else {
			fmt.Printf("greate guess you won in %v guesses\n", i)
			break
		}

	}

}

func main() {
	fmt.Println("Random")
	randomNumber()
}
