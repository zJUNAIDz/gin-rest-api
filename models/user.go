package models

import (
	"errors"

	"github.com/zjunaidz/gin-rest-api/db"
	"github.com/zjunaidz/gin-rest-api/utils"
)

type User struct {
	Id       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (user User) Save() error {
	query := `
	INSERT INTO users (email, password)
	VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	//MS used separate variable to store it but I feel like this is better approach to eliminate any plain password residue
	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(user.Email, user.Password)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.Id = userId
	return nil
}

func (user User) ValidateCredentials() error {
	query := `
	SELECT password
	FROM users
	WHERE email = ?
	`
	row := db.DB.QueryRow(query, user.Email)
	if row.Err() != nil {
		return row.Err()
	}
	var retrievedHashedPassword string
	err := row.Scan(&retrievedHashedPassword)
	if err != nil {
		return err
	}
	isPasswordCorrect := utils.CompareHashedPassword(user.Password, retrievedHashedPassword)
	if !isPasswordCorrect {
		return errors.New("credentials invalid. Incorrect password")
	}
	return nil
}

func GetAllUsers() ([]User, error) {
	query := `
	SELECT * FROM users
	`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//MS: I think it's better to use pointer to the slice of users
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
