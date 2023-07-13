package auth

import (
	"crud-simple-api/db"
	"crud-simple-api/models"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GenerateJWT(userID uint64, email string, secretKey string, expiration time.Duration) (string, error) {

	claims := &models.Claims{
		UserID: userID,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func AuthenticateUser(email, password string) (*models.User, error) {

	user := &models.User{}

	// Fetch the user from database on the provided email
	if err := db.DB.Where("email = ?", email).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}

		return nil, err
	}

	// compare hased password with the provided password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return nil, errors.New("Invalid password")
		}

		return nil, err
	}

	return user, nil
}