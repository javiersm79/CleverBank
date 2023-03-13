package dto

import "github.com/jinzhu/gorm"

type MovementAccount struct {
	gorm.Model
	MovementId           string `gorm:"unique"`
	Action               string
	SourceAccountNumber  string
	DestinyAccountNumber string
	Amount               int64
}
