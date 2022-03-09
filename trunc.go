package main

/*
Write a program which prompts the user to enter a floating point number and prints the integer which is a truncated version of the floating point number that was entered. Truncation is the process of removing the digits to the right of the decimal place.

Submit your source code for the program, “trunc.go”.

*/

import "fmt"

func truncate(num float64) int {
	roundNum := int(num)
	return roundNum
}

func main() {
	var userInput int
	fmt.Println("Give me a decimal number. I'm going to round it for you. ")
	fmt.Scan(&userInput)

	fmt.Println((float64(userInput)))
	return
}
