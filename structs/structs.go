package main

import (
	"fmt"
	//"os/user"

	"example.com/structs/user"
)

func main() {
	UserfirstName := user.GetUserData("Please enter your first name: ")
	UserlastName := user.GetUserData("Please enter your last name: ")
	Userbirthdate := user.GetUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User

	appUser, err := user.New(UserfirstName, UserlastName, Userbirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	// ... do something awesome with that gathered data!

	admin := user.NewAdmin("test", "test123")

	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	//fmt.Println(firstName, lastName, birthdate)
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}
