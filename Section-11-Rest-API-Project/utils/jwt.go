package utils

import "github.com/golang-jwt/jwt/v5"
import "time"


const secretkey = "mysecretkey"

func GenerateToken(email string, userId int64) (string, error){
	// Generate a random token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"expiry": time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretkey))
}
