package main

import (
	"fmt"
	"example.com/structs/userStruct"
)


func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// var appUser User

	// struct literal
	/*appUser := User{
		firstName: userFirstName,
		lastName:  userLastName,
		birthdate: userBirthdate,
		craetedAt: time.Now(),
	}

		also write like this, make sure that order is same
		appUser = User{
			userFirstName ,
			userLastName,
			userBirthdate,
			time.Now(),
		}*/

	// also write using constructor
	var appUser *userStruct.User
	appUser, err := userStruct.NewUser(userFirstName, userLastName, userBirthdate)
		if err != nil {
			fmt.Println(err)
			return
		}

	appUser.OutputUserData()
	appUser.ClearUserName()
	appUser.OutputUserData()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
