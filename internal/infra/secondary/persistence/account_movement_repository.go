package persistence

import (
	Account "cleverbank/internal/core/domain/account"
	"cleverbank/internal/infra/secondary/mapper"
	"cleverbank/internal/infra/secondary/persistence/dto"
	"fmt"
	"github.com/jinzhu/gorm"
)

type AccountMovementRepository struct {
	Dbconex *gorm.DB
}

func (air AccountMovementRepository) Deposit(request Account.AccountMovementRequest) (Account.AccountMovementResponse, error) {
	var dbAccount dto.Account

	if err := air.Dbconex.Where("number = ?", request.DestinyAccountNumber).First(&dbAccount).Error; err != nil {
		return Account.AccountMovementResponse{}, fmt.Errorf("Account not found: %s", request.DestinyAccountNumber)
	}

	// Update with conditions
	air.Dbconex.Model(&dto.Account{}).Where("number = ?", request.DestinyAccountNumber).Update("balance", dbAccount.Balance+request.Amount)

	accountMovementResponse := mapper.MapAccountMovementDtoToDomain(dbAccount, request)
	return accountMovementResponse, nil
}
