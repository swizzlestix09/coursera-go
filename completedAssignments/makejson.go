// Write a program which prompts the user to first enter a name, and then enter an address. Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively. Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.

// Submit your source code for the program,
// “makejson.go”.

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func createMap(userName string, userAddress string) map[string]string {
	userInfo := make(map[string]string)
	userInfo["name"] = userName
	userInfo["address"] = userAddress
	return userInfo
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Provide a name: ")
	scanner.Scan()
	name := scanner.Text()
	fmt.Println("Provide an address: ")
	scanner.Scan()
	address := scanner.Text()

	var userInfo = createMap(name, address)

	barr, err := json.Marshal(userInfo)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(barr))
	return

}
