package userStruct

import (
	"fmt"
	"time"
	"errors"
)


type User struct {
	firstName string
	lastName  string
	birthdate string
	craetedAt time.Time

}

type Admin struct {
	email string
	password string
	UserDetails User
}

// function in struct where it is used as a receiver
func (u *User) OutputUserData() {
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.craetedAt)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// constructor function
func NewUser(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("empty fields are invalid")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		craetedAt: time.Now(),
	}, nil
} // you can also return a pointer so that no copy is created

func NewAdmin(email, password string) Admin {
	return Admin {
		email: email, 
		password: password, 
		UserDetails: User {
			firstName: "ADMIN",
			lastName: "Admin",
			birthdate: "----",
			craetedAt: time.Now(),
		},
	}
}