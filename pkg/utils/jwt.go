package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const SECRET_KEY string = "*****************"

var EXPIRATION_TIME int64 = time.Now().Add(time.Hour * 2).Unix()

// generate a JWT Token signed
func GenerateToken(email string, userID int64) (string, error) {
	// signin method and user data in maps, expiration
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userID": userID,
		"exp":    EXPIRATION_TIME,
	})
	// return a pointer to the
	// send to client as single string and error incase
	return token.SignedString([]byte(SECRET_KEY))
}
