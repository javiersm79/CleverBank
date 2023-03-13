package usecase

import Account "cleverbank/internal/core/domain/account"

type AccountInfoUseCase interface {
	Handle(accountNumber string) (Account.AccountDetails, error)
}
