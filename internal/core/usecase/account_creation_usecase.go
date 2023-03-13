package usecase

import Account "cleverbank/internal/core/domain/account"

type AccountCreationUseCase interface {
	Handle(newAccountReq Account.NewAccountReq) (Account.AccountDetails, error)
}
