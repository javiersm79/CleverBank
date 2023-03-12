package persistence

import (
	Account "cleverbank/internal/core/domain/account"
	"cleverbank/internal/infra/secondary/mapper"
	"cleverbank/internal/infra/secondary/persistence/dto"
	"fmt"
	"github.com/jinzhu/gorm"
)

type AccountInfoRepository struct {
	Dbconex *gorm.DB
}

func (air AccountInfoRepository) GetBalance(accountNumber string) (Account.AccountDetails, error) {
	var dbAccount dto.Account

	if err := air.Dbconex.Where("number = ?", accountNumber).First(&dbAccount).Error; err != nil {
		return Account.AccountDetails{}, fmt.Errorf("Account not found: %s", accountNumber)
	}

	acountDetails := mapper.MapAccountInfoDtoToDomain(dbAccount)
	return acountDetails, nil
}
