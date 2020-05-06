/*
Example usage of bcrypt package.
*/

package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

//Function to create a hash using given password string
//Inputs: password(string)
//Outputs: hash([]byte) and error(if any else nil)
func hashPassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		fmt.Println("Error creating hash of password")
		return []byte{}, err
	}
	return hash, err
}

//Function to compare stored hash and new password string
//Inputs: password(string) and stored hash([]byte)
//Outputs: status(bool)
func comparePasswordWithHash(password string, hash_stored []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash_stored, []byte(password))
	if err != nil {
		return false
	}
	return true
}

//Function that sends response about login status
//Inputs: status(bool)
//Ouputs: None
func sendResponse(status bool) {
	if status {
		fmt.Println("Login successful\n\n")
	} else {
		fmt.Println("Login failed\n\n")
	}
}

//Example
func main() {
	//Initial hashing and hashed value is stored
	password := "abcdef1234"
	hashed_password, err := hashPassword(password)
	if err != nil {
		fmt.Println("Unable to hash the password. Exiting....\n")
		return
	}
	fmt.Println("Password hashed and stored successfully\n")

	//Example of failed login
	password_entered := "abcde12"
	fmt.Println("Received a login request. Validating password....")
	sendResponse(comparePasswordWithHash(password_entered, hashed_password))

	//Example of failed login
	password_entered = "abcdef1234"
	fmt.Println("Received a login request. Validating password....")
	sendResponse(comparePasswordWithHash(password_entered, hashed_password))
}

/*
Output:
Password hashed and stored successfully

Received a login request. Validating password....
Login failed


Received a login request. Validating password....
Login successful
*/
