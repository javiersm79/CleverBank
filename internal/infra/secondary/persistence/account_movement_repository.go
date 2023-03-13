package persistence

import (
	Account "cleverbank/internal/core/domain/account"
	"cleverbank/internal/infra/secondary/mapper"
	"cleverbank/internal/infra/secondary/persistence/dto"
	"fmt"
	"github.com/google/uuid"
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

	air.Dbconex.Model(&dto.Account{}).Where("number = ?", request.DestinyAccountNumber).Update("balance", dbAccount.Balance+request.Amount)
	dbMovementAccount := GenerateMovement(request)
	accountMovementResponse := mapper.MapAccountMovementDtoToDomain(dbMovementAccount)

	result := air.Dbconex.Create(&dbMovementAccount) // pass pointer of data to Create

	if result.Error != nil {
		return Account.AccountMovementResponse{}, fmt.Errorf("Error generating movement: %s", result.Error.Error())
	}

	return accountMovementResponse, nil
}

func GenerateMovement(request Account.AccountMovementRequest) dto.MovementAccount {
	dbMovementAccount := dto.MovementAccount{
		Model:                gorm.Model{},
		MovementId:           uuid.NewString(),
		Action:               request.Type,
		SourceAccountNumber:  request.SourceAccountNumber,
		DestinyAccountNumber: request.DestinyAccountNumber,
		Amount:               request.Amount,
	}
	return dbMovementAccount
}
