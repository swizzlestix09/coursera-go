package main

/*
Write a program which prompts the user to enter a string. The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’. The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’. The program should print “Not Found!” otherwise. The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

Examples: The program should print “Found!” for the following example entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”. The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.

*/

//string has to start with letter i
// end with the letter n
//somewhere in between has to be an a

import "fmt"

func ianVerfification(name string) string {
	var icheck string = "i"
	var Icheck string = "I"
	var ncheck string = "n"
	var Ncheck string = "N"

	var acheck string = "a"
	var Acheck string = "A"
	var firstLet string = string(name[0])
	var lastLet string = string(name[len(name)-1])

	var fd string = "Found!"
	var notFd string = "Not Found!"

	if firstLet != icheck && firstLet != Icheck {
		return notFd
	}
	if lastLet != ncheck && lastLet != Ncheck {
		return notFd
	}

	for i := 1; i < len(name)-2; i++ {
		if string(name[i]) == acheck || string(name[i]) == Acheck {
			return fd
		}
	}

	return notFd
}

func main() {
	fmt.Println("Give me a piece of your soul...JK. A word will do: ")
	var wordGiven string
	fmt.Scan(&wordGiven)

	fmt.Println(ianVerfification(wordGiven))

}

// verified := "Found!"
// unverified := "Not Found!"

// 	if firstLet != 'i' ||firstLet != 'I' {
// 		return false
// 	}
