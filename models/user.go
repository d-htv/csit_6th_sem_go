package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct{
	ID uint64 `gorm:"autoIncrement;not null;primaryKey;" json:"id"`
	UserDto
	CreatedAt time.Time `gorm:"autoCreateTime;" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type UserDto struct{
	FirstName string `json:"first_name"`
	LastName string `json:"second_name"`
	Email string `json:"email"`
	Contact string `json:"contact"`
	Age uint8 `gorm:"age"`
}