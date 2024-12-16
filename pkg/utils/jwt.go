package utils

import (
	"apiv2/pkg/config"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SECRET_KEY string = config.Envs.SECRET_KEY
var TOKEN_EXPIRATION_HOURS = time.Duration(config.Envs.EXPIRATION_TIME)

// TokenBlacklist to store invalidated tokens
var tokenBlacklist = make(map[string]bool)
var blacklistMutex = &sync.Mutex{}

// GenerateToken generates a JWT token with a dynamic expiration time
func GenerateToken(email string, userID int64) (string, error) {
	expirationTime := time.Now().Add(time.Hour * TOKEN_EXPIRATION_HOURS).Unix()

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userID": userID,
		"iss":    "eventsAPI",
		"exp":    expirationTime,
		"iat":    time.Now().Unix(),
	})

	// Sign the token and return it
	return token.SignedString([]byte(SECRET_KEY))
}

// VerifyToken checks the validity of the token and returns the userID if valid
func VerifyToken(token string) (float64, error) {
	// Check if the token is blacklisted
	blacklistMutex.Lock()
	if tokenBlacklist[token] {
		blacklistMutex.Unlock()
		return 0, errors.New("token has been logged out")
	}
	blacklistMutex.Unlock()

	// Parse the token
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is as expected
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method in JWT token")
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return 0, errors.New("token verification failed: " + err.Error())
	}

	if !parsedToken.Valid {
		return 0, errors.New("invalid or expired token")
	}

	// Extract claims and assert type
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("mismatched token claims")
	}

	// Extract userID safely
	userID, ok := claims["userID"].(float64)
	if !ok {
		return 0, errors.New("invalid userID in token claims")
	}

	return userID, nil
}

// Logout invalidates a token by adding it to the blacklist
func Logout(token string) {
	blacklistMutex.Lock()
	defer blacklistMutex.Unlock()
	tokenBlacklist[token] = true
	fmt.Println("Token successfully logged out.")
}
