package models

import (
	"database/sql"
	"errors"
	"fmt"
	"golang-crud-app/db"
	"golang-crud-app/utils"
)

type User struct {
	ID       int64  `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Println("Could not close stmt.")
		}
	}(stmt)

	hashedPasswd, err := utils.HashPasswd(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPasswd)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = userId

	return nil
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	var retrievedPassword string

	err := db.DB.QueryRow(query, u.Email).Scan(&u.ID, &retrievedPassword)
	if err != nil {
		fmt.Println(err)
		return errors.New("Invalid credentials : unable to retrieve credentials from database")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("Invalid credentials : password is incorrect")
	}

	return nil
}
