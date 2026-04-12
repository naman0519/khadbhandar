package models

import "time"

type Order struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `form:"name"`
	Product   string `form:"product"`
	Phone     string `form:"phone"`
	Quantity  int    `form:"quantity"`
	Status    string `form:"status"`
	CreatedAT time.Time
}
