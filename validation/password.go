package validation

import (
	"crud-simple-api/helper"

	"golang.org/x/crypto/bcrypt"
)

// Check if the provided password meet the requirements
func ValidatePassword(password string) bool {

	if len(password) < 8 {
		return false
	}

	return true

}

// Generate Bcrypt hash for the password
func EncryptPassword(password string) (string, error) {
	
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPassword), nil
	
}

// Compare the provided password with the stored password
func ComparePassword(password, hashedPassword string) error {
	
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	helper.PanicIfError(err)
	
	return nil

}