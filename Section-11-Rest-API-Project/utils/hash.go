// file responsible for hashing passwords

package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error){
	// Use the bcrypt library to hash the password
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password string, hashedPassword string) bool {
	// Use the bcrypt library to check the password against the hash
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
