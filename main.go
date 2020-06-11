package main

import (
	"fmt"
	domain "github.com/sonatard/sample-project/user"
)

func main() {
	user := domain.NewUser("foo", "bar")
	fmt.Printf("FirstName: %v\n", user.FirstName)
	fmt.Printf("LastName: %v\n", user.LastName)

	// remove following comment out, but warning is not deleted.
	// fmt.Printf("Name: %v\n", user.Name())
}
