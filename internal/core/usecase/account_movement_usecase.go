package usecase

import "cleverbank/internal/core/domain/account"

type AccountMovementUseCase interface {
	Handle(movementRequest account.AccountMovementRequest) (account.AccountMovementResponse, error)
}
