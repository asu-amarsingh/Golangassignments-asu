// Author: Amar Singh
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	var items []string
	//loop to prompt user and ask for at *least* 3 items or more
	fmt.Println("Hello user, I'm going to have to ask you to input several (3 or more) strings")

	for i := 0; i <= len(items); i++ {

		//add string prompt
		fmt.Print("String to add: ")

		//take in user input
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		//remove extra space
		input = strings.TrimSpace(input)

		//append (store the input into the slice)
		items = append(items, input)

		//ask to continue
		var yesNo string
		fmt.Print("Continue? [Y/N]:")
		fmt.Scanln(&yesNo)

		//if y or Y continue, if n or N break ( unless the user didn't input 3 items!)
		if yesNo == "y" || yesNo == "Y" {
			continue
		} else if i < 2 {
			fmt.Println("That is NOT 3 or more strings >:(")

		} else if yesNo == "n" || yesNo == "N" {
			break
		}
	}
	fmt.Println(strings.Join(items, ", "))
}
