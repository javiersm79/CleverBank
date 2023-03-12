package mapper

import (
	"cleverbank/internal/core/domain/account"
	"cleverbank/internal/infra/secondary/persistence/dto"
)

func MapAccountInfoDtoToDomain(accountDto dto.Account) account.AccountDetails {
	accountDetails := account.AccountDetails{
		Id:        int64(accountDto.ID),
		CreatedAt: accountDto.CreatedAt,
		UpdatedAt: accountDto.UpdatedAt,
		DeletedAt: accountDto.DeletedAt,
		ClientId:  accountDto.ClientId,
		Type:      accountDto.Type,
		Number:    accountDto.Number,
		Balance:   accountDto.Balance,
	}
	return accountDetails
}
