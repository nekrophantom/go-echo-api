package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	UserID 	uint64 	`json:"userId"`
	Email	string	`json:"email"`
	jwt.StandardClaims
}