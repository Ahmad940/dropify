package cli

import (
	"fmt"

	"github.com/Ahmad940/dropify/app/model"
	"github.com/Ahmad940/dropify/app/service"
)

// NewUser create a new user
func NewUser(username, password, role *string) {
	fmt.Println("Creating new user...")
	err := service.CreateAccount(model.Auth{
		UserName: *username,
		Password: *password,
		Role:     *role,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("User created successful")
}

// ResetUser reset a user password
func ResetUser(username, password *string) {
	fmt.Println("Updating user...")
	err := service.UpdateUserPassword(*username, *password)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Updated user")
}

// DeleteUser delete a user
func DeleteUser(username *string) {
	fmt.Println("Creating new user...")
	err := service.DeleteUser(*username)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Deleted successful")
}
