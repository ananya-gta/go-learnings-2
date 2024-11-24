// file responsible for hashing passwords

package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error){
	// Use the bcrypt library to hash the password
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
