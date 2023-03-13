package port

import "cleverbank/internal/core/domain/account"

type AccountMovementPort interface {
	Deposit(movementRequest account.AccountMovementRequest) (account.AccountMovementResponse, error)
	Transfer(movementRequest account.AccountMovementRequest) (account.AccountMovementResponse, error)
	Withdrawal(movementRequest account.AccountMovementRequest) (account.AccountMovementResponse, error)
}
