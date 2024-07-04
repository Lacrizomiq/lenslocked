// Messing around with bcrypt

package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	switch os.Args[1] {
	case "hash":
		hash(os.Args[2])
	case "compare":
		compare(os.Args[2], os.Args[3])
	default:
		fmt.Printf("Invalid command: %v\n", os.Args[1])
	}

}

func hash(password string) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Hashed password: %s\n", string(hashBytes))
}

func compare(password, hash string) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		fmt.Printf("Password is invalid: %v\n", password)
		return
	}
	fmt.Println("Password is correct")
}
