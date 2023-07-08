package db

import (
	"crud-simple-api/helper"
	"crud-simple-api/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{})
	helper.PanicIfError(err)

	return nil
}