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

// function in struct where it is used as a receiver
func (u *User) outputUserData() {
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.craetedAt)
}

func (u *User) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// constructor function
func newUser(firstName, lastName, birthdate string) *User {
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		craetedAt: time.Now(),
	}
} // you can also return a pointer so that no copy is created

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
	appUser := *newUser(userFirstName,
		userLastName,
		userBirthdate)

	appUser.outputUserData()
	appUser.clearUserName()
	appUser.outputUserData()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
