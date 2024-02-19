package models

import (
	"gorm.io/gorm"
)

type Person struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

type User struct {
	Name     string `json:"name" validate:"required,min=3,max=32"`
	IsActive *bool  `json:"isactive" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required,email,min=3,max=32"`
}

type Company struct {
	Email        string `json:"email,omitempty" validate:"required,email" regexp:"^[a-zA-Z0-9-]{2,30}$"`
	Username     string `json:"username,omitempty" validate:"required"`
	Password     string `json:"password,omitempty" validate:"required,min=6,max=20"`
	LineID       string `json:"lineId" validate:"required"`
	PhoneNumber  string `json:"phoneNumber" validate:"required"`
	BusinessType string `json:"businessType" validate:"required"`
	WebsiteName  string `json:"websiteName" validate:"required" regexp:"^[a-z0-9-]{2,30}$"`
}

type Dogs struct {
	gorm.Model
	DogID int    `json:"dog_id"`
	Name  string `json:"name"`
}
