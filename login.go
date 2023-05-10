// Author: Amar Singh
package main

import (
	"fmt"
)

func main() {
	//variables
	loginID := "admin"
	password := "Pa$$w0rd"
	attempts := 3

	//for loop with 3 attempts
	for attempts > 0 {

		//prompt login
		fmt.Print("Please Enter the Login ID: ")
		var inputID string
		fmt.Scanln(&inputID)

		//prompt password
		fmt.Print("Enter the Password: ")
		var inputPassword string
		fmt.Scanln(&inputPassword)

		//check credentials
		if inputID == loginID && inputPassword == password {
			fmt.Println("Login Successful!")
			break
		} else if inputID == loginID {
			attempts--
			fmt.Printf("Password Incorrect! Try again - %d attempts left!\n", attempts)
		} else if inputPassword == password {
			attempts--
			fmt.Printf("Login Incorrect! Try again - %d attempts left!\n", attempts)
		} else {
			attempts--
			fmt.Printf("Login and Password Incorrect! Try again - %d attempts left!\n", attempts)
		}
	}

	//account lock
	if attempts == 0 {
		fmt.Println("Account locked – Contact 800-123-4567.")
	}
}
