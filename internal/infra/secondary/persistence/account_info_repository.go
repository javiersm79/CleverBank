package persistence

import (
	"cleverbank/internal/infra/secondary/persistence/dto"
	"github.com/jinzhu/gorm"
)

type AccountInfoRepository struct {
	Dbconex *gorm.DB
}

func (air AccountInfoRepository) CreateAccount(account string) (string, error) {
	var dbuser dto.User
	air.Dbconex.Where("email = ?", "user@test.com").First(&dbuser)
	return dbuser.Password, nil
}

func (air AccountInfoRepository) GetBalance(accountNumber string) (string, error) {
	var dbuser dto.User
	air.Dbconex.Where("email = ?", "user@test.com").First(&dbuser)
	return dbuser.Password, nil
}
