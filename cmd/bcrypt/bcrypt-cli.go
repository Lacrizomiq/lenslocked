// Messing around with bcrypt

package main

import (
	"fmt"
	"os"

	"crypto/sha256"

	"encoding/hex"
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
	fmt.Printf("Todo: Hash the password %q\n", password)

}

func compare(password, hash string) {
	fmt.Printf("Todo: Compare the password %q with the hash %q\n", password, hash)
}
