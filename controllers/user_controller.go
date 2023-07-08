package controllers

import (
	"crud-simple-api/db"
	"crud-simple-api/helper"
	"crud-simple-api/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)


func GetAllUsers(c echo.Context) error {

	var users []models.User

	result := db.DB.Find(&users)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.Response(http.StatusInternalServerError, "Error retrieving users", nil))
	}

	response := helper.Response(http.StatusOK, "Users retrieved successfully", users)
	return c.JSON(http.StatusOK, response)
}

func CreateUser(c echo.Context) error {

	var user models.User

	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.Response(http.StatusBadRequest, "Invalid request payload", nil))
	}

	// Encrypt Password with Bcrypt
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Response(http.StatusInternalServerError, "Error encrypting password", nil))
	}
	user.Password = string(hashPassword)

	result := db.DB.Create(&user)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.Response(http.StatusInternalServerError, "Error creating user", nil))
	}

	response := helper.Response(http.StatusOK, "User create successfully", user)
	return c.JSON(http.StatusCreated, response)
}

func GetUserById(c echo.Context) error {

	// Get Id from Params
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.Response(http.StatusBadRequest, "User not found", nil))
	}
	var user models.User
	
	// Search model User by id
	result := db.DB.First(&user, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.Response(http.StatusInternalServerError, "Error retrieving user", nil))
	}

	// Output response
	response := helper.Response(http.StatusOK, "User retrieved successfully", user)
	return c.JSON(http.StatusOK, response)
}

func UpdateUser(c echo.Context) error {
	
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.Response(http.StatusBadRequest, "User not found", nil))
	}
	var user models.User

	// Search model User by id
	result := db.DB.First(&user, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.Response(http.StatusInternalServerError, "User not found", nil))
	}

	// Check request payload
	err = c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.Response(http.StatusBadRequest, "Invalid request payload", nil))
	}

	// Check if the password field is provided and encrypt it
	if user.Password != ""{
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.Response(http.StatusInternalServerError, "Error encrypting password", nil))
		}
		user.Password = string(hashPassword)
	}

	// Save update user
	result = db.DB.Save(&user)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.Response(http.StatusInternalServerError, "Error updating user", nil))
	}

	response := helper.Response(http.StatusOK, "User updated successfully", user)
	return c.JSON(http.StatusOK, response)
}

func DeleteUser(c echo.Context) error {
	
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.Response(http.StatusBadRequest, "User not found", nil))
	}
	var user models.User

	result := db.DB.First(&user, id)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, helper.Response(http.StatusInternalServerError, "User not found", nil))
	}

	result = db.DB.Delete(&user)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, helper.Response(http.StatusInternalServerError, "Error deleting user", nil))
	}

	response := helper.Response(http.StatusOK, "User deleted successfully", nil)
	return c.JSON(http.StatusOK, response)
}