package models

import "time"

type Order struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	createdAt time.Time
	ProductID uint    `json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID"`
	UserID    uint    `json:"user_id"`
	User      User    `gorm:"foreignKey:UserID"`
}
