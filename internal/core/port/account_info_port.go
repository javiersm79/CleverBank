package port

import Account "cleverbank/internal/core/domain/account"

type AccountInfoPort interface {
	GetBalance(accountNumber string) (Account.AccountDetails, error)
}
