package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	craetedAt time.Time
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// var appUser User

	// struct literal
	appUser := User{
		firstName: userFirstName,
		lastName:  userLastName,
		birthdate: userBirthdate,
		craetedAt: time.Now(),
	}
	/*
		also write like this, make sure that order is same
		appUser = User{
			userFirstName ,
			userLastName,
			userBirthdate,
			time.Now(),
		}*/

	outputUserData(appUser)

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}

func outputUserData(u User) {
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.craetedAt)
}
