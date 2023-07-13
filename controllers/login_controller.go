package controllers

import (
	"crud-simple-api/auth"
	"crud-simple-api/config"
	"crud-simple-api/helper"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	// AuthController fields
}

func NewAuthController() *AuthController {
	return &AuthController{
		// Initialize AuthController fields if any
	}
}

func Login(c echo.Context) error {

	email := c.FormValue("email")
	password := c.FormValue("password")

	// Authenticate user
	user, err := auth.AuthenticateUser(email, password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.Response(http.StatusUnauthorized, "Invalid credentials", nil))
	}

	// Generate JWT Token
	token, err := auth.GenerateJWT(uint64(user.Id), user.Email, config.LoadConfig().JWTSecret, config.LoadConfig().JWTExpiration)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Response(http.StatusInternalServerError, "Failed to generate token", nil))
	}

	cookie := &http.Cookie{
		Name: "token",
		Value: token,
		Expires: time.Now().Add(config.LoadConfig().JWTExpiration),
		Path: "/",
	}

	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "Login successfull",
		"token" : token,
	})
}