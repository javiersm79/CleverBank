package dto

import "github.com/jinzhu/gorm"

type Account struct {
	gorm.Model
	ClientId int64
	number   int64
	Type     string
	balance  float64
	//Email    string `gorm:"unique"`
}
