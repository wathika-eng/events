package utils

import (
	"errors"
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

// verify if token from client is legit
func VerifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Check if the token was signed using the expected method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method in JWT token")
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return errors.New("token verification failed: " + err.Error())
	}

	if !parsedToken.Valid {
		return errors.New("invalid or expired token")
	}

	// Extract claims and assert type
	_, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("mismatched token claims")
	}

	return nil
}
