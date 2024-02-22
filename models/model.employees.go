// +build ignore
package models

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	EmployeeID int        `json:"employee_id" gorm:"primaryKey"`
	Name       string     `json:"name"`
	LastName   string     `json:"last_name"`
	Birthday   *time.Time `json:"birthday"`
	Age        *int       `json:"age"`
	Email      string     `json:"email"`
	Tel        *string    `json:"tel"`
}
