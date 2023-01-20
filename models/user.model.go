package models

import (
	db "github.com/wpcodevo/golang-postgresql-api/db/sqlc"
)

type SignInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	UserId       	string 		`json:"id,omitempty"`
	Email			string 		`json:"email,omitempty"`
	Gender			string 		`json:"gender,omitempty"`
	PhoneNumber		string 		`json:"phonenumber,omitempty"`
	PasswordHash	string		`json:"lastname,omitempty"`
	LastName		string 		`json:"lastname,omitempty"`
	FirstName		string 		`json:"firstname,omitempty"`
}
