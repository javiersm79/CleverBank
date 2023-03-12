package dto

import "github.com/jinzhu/gorm"

type Account struct {
	gorm.Model
	ClientId int64
	Number   string `gorm:"unique"`
	Type     string
	Balance  int64
}
