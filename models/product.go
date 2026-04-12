package models

type Product struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Category string
	Price    string
	Image    string
	Stock    int
}
