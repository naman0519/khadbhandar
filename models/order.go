package models

import "time"

type Order struct {
	ID          uint   `gorm:"primaryKey"`
	UserName    string `form:"name"`
	Product     string `form:"product"`
	PhoneNumber string `form:"phone"`
	Quantity    int    `form:"quantity"`
	Status      string `form:"status"`
	CreatedAt   time.Time
}
