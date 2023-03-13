package port

import Account "cleverbank/internal/core/domain/account"

type AccountCreationPort interface {
	CreateAccount(newAccountRequest Account.NewAccountReq) (Account.AccountDetails, error)
}
