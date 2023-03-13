package dto

import "github.com/jinzhu/gorm"

type MovementAccount struct {
	gorm.Model
	MovementId           string `gorm:"unique"`
	Action               string
	Secuence             string
	SourceAccountNumber  string
	DestinyAccountNumber string
	Amount               int64
}
