package dto

import "github.com/jinzhu/gorm"

type Client struct {
	gorm.Model
	Name     string
	LastName string
	Rut      string `gorm:"unique"`
	Email    string `gorm:"unique"`
}
