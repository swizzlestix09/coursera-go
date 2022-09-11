// Write a program which reads information from a file and represents it in a slice of structs. Assume that there is a text file which contains a series of names. Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

// Your program will define a name struct which has two fields, fname for the first name, and lname for the last name. Each field will be a string of size 20 (characters).

// Your program should prompt the user for the name of the text file. Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file. Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file. After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.

// Submit your source code for the program, “read.go”.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var MAX_LENGTH = 20

//define name struct - name lastname, string of size 20
type Fullname struct {
	fName string `valid:MaxSize(20)`
	lName string `valid:MaxSize(20)`
}

func makeStruct(words []string) Fullname {
	var name Fullname
	name.fName = words[0]
	name.lName = words[1]
	return name
}

func main() {
	//prompt user for name of text file
	var nameOfFile string
	names := []Fullname{
		Fullname{
			fName: "angie",
			lName: "rodriguez",
		},
	}
	fmt.Println("Name of file please..")
	fmt.Scan(&nameOfFile)

	//read ea line
	readFile, err := os.Open(nameOfFile)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	//name added to struct
	for fileScanner.Scan() {
		line := fileScanner.Text()
		//struct added to slice

		words := strings.Fields(line)

		nameStruct := makeStruct(words)
		names = append(names, nameStruct)
	}

	//iterate through slice, printing first and last names
	for _, name := range names {
		fmt.Println(name.fName, name.lName)
	}
	return

}
