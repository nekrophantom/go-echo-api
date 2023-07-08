package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id			uint			`gorm:"primaryKey" json:"id"`
	Name		string			`gorm:"type:varchar(255)" json:"name"`
	Email		string			`gorm:"type:varchar(255)" json:"email"`
	Password	string			`gorm:"type:varchar(255)" json:"password"`
	CreatedAt	time.Time		`gorm:"index" json:"created_at"`
	UpdatedAt	time.Time		`gorm:"index" json:"updated_at"`
	DeletedAt	gorm.DeletedAt	`gorm:"index" json:"deleted_at,omitempty"`
}