package middleware

import (
	"crud-simple-api/config"
	"crud-simple-api/helper"
	"crud-simple-api/models"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	
	return func(c echo.Context) error {

		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, helper.Response(http.StatusUnauthorized, "Unauthorized", nil))
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.LoadConfig().JWTSecret), nil
		})

		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.Response(http.StatusUnauthorized, "Invalid Token", nil))
		}

		claims, ok := token.Claims.(*models.Claims)
		if !ok || !token.Valid {
			return c.JSON(http.StatusUnauthorized, helper.Response(http.StatusUnauthorized, "Invalid Token", nil))
		}

		c.Set("userId", claims.UserID)
		c.Set("email", claims.Email)

		return next(c)
	}

}