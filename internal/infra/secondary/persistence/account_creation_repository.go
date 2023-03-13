package persistence

import (
	Account "cleverbank/internal/core/domain/account"
	"cleverbank/internal/infra/secondary/accountType"
	"cleverbank/internal/infra/secondary/persistence/dto"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type AccountCreationRepository struct {
	Dbconex *gorm.DB
}

func (acr AccountCreationRepository) CreateAccount(newAccountRequest Account.NewAccountReq) (Account.AccountDetails, error) {
	var dbAccount dto.Account

	//acountDetails := mapper.MapAccountInfoDtoToDomain(dbAccount)*/
	acountDetails, err := accountType.GetAccount(newAccountRequest.Type)

	if err != nil {
		return Account.AccountDetails{}, fmt.Errorf("Error Creating new Account: %s", err.Error())
	}

	resultFind := acr.Dbconex.Where("number = ?", acountDetails.GetNumberAccount()).First(&dbAccount)

	if resultFind.RowsAffected > 0 {
		return Account.AccountDetails{}, fmt.Errorf("Account  found: %s", dbAccount.Number)
	}
	dbAccount.Number = acountDetails.GetNumberAccount()
	dbAccount.ClientId = newAccountRequest.ClientId
	dbAccount.Type = newAccountRequest.Type
	resultCreate := acr.Dbconex.Create(&dbAccount)

	if resultCreate.Error != nil {
		return Account.AccountDetails{}, fmt.Errorf("Error Creating new Account: %s", resultCreate.Error.Error())
	}

	return Account.AccountDetails{
		Id:        int64(dbAccount.ID),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
		ClientId:  newAccountRequest.ClientId,
		Type:      dbAccount.Type,
		Number:    dbAccount.Number,
		Balance:   0,
	}, nil
}
