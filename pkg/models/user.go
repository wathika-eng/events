// package model defines interfaces and methods for the DB
package models

import (
	"apiv2/pkg/db"
	"apiv2/pkg/utils"
	"database/sql"
	"errors"
	"fmt"
)

// shape of user
type User struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	CreatedAt string `json:"time"`
}

// save user to Database
func (u *User) Save() error {
	// Set the creation time
	u.CreatedAt = utils.DBTime()

	// Hash the user's password
	hashedPass, err := utils.HashPassword(u.Password)
	if err != nil {
		return fmt.Errorf("error hashing password: %v", err)
	}

	// Define the query to insert the user (prevents SQL injection)
	query := `
		INSERT INTO users (name, email, password, time)
		VALUES (?, ?, ?, ?)
	`

	// Execute the query directly
	result, err := db.DB.Exec(query, u.Name, u.Email, hashedPass, u.CreatedAt)
	if err != nil {
		return fmt.Errorf("error executing the SQL query: %v", err)
	}

	// Get the last inserted ID
	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error retrieving last inserted ID: %v", err)
	}

	// Assign the new ID to the user struct
	u.ID = id

	return nil
}

// validate credentials
func (u *User) ValidateCreds() error {
	const query string = `SELECT id, password FROM USERS WHERE email = ?`
	var retrievedPass string

	// Query the database for the user's password
	err := db.DB.QueryRow(query, u.Email).Scan(&u.ID, &retrievedPass)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("email not found")
		}
		return fmt.Errorf("database error: %w", err)
	}

	// Check if the password matches
	if !utils.CheckPass(u.Password, retrievedPass) {
		return errors.New("wrong password")
	}

	return nil
}
