// Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
// The program should be written as a loop.
// Before entering the loop, the program should create an empty integer slice of size (length) 3.
// During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
// The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
// The slice must grow in size to accommodate any number of integers which the user decides to enter.
// The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.

// Submit your source code for the program,
// “slice.go”.
package main

import (
	"fmt"
)

//make function that takes array and number and sorts.
func organizeNums() {
	//sort
}

func main() {
	//declare a empty  slice that is size 3.
	var userInput int
	inputInts := make([]int, 3)

	for i := 0; i < 2; i++ {
		fmt.Println("number please..")
		fmt.Scan(&userInput)
		fmt.Println(inputInts, userInput)
	}

	//prompt user to enter intergers
	//store intergers in sorted slice

	fmt.Println("Done picking numbers. Here they are, sorted by the computer because you're a lazy homosapien: ", inputInts)
	//if user types X quit program
}
