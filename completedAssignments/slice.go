// Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
// The program should be written as a loop.
// Before entering the loop, the program should create an empty integer slice of size (length) 3.
// During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
// The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
// The slice must grow in size to accommodate any number of integers which the user decides to enter.
// The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.

package main

import (
	"fmt"
	"sort"
	"strconv"
)

//make function that takes array and sorts.
func organizeNums(array []int) []int {
	sort.Ints(array)
	fmt.Println(array)
	return array
}

func main() {
	var userInput string
	inputInts := make([]int, 3)

	for userInput != "X" || userInput != "x" {

		fmt.Println("number please..")g
		fmt.Scan(&userInput)
		num, err := strconv.Atoi(userInput)
		if err != nil {
			//there is no notes on how to handle non "X" letters.
			//Since previously they were being made to 0 before strconv, I kept it that way.
			if userInput == "X" || userInput == "x" {
				fmt.Println("Oh, good. Maybe you'll get off your phone and do it yourself.")
				return
			}
			inputInts = append(inputInts, 0)
		}

		inputInts = append(inputInts, num)
		inputInts = organizeNums(inputInts)
	}

	fmt.Println("Done picking numbers. Here they are, sorted by my code because you're a lazy homosapien: ", inputInts)
}
