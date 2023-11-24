package models

import "time"

type User struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name"`
}
