package instance

import (
	"cleverbank/internal/core/implementation"
	"cleverbank/internal/core/usecase"
)

func GetAccountMovementUseCase() usecase.AccountMovementUseCase {
	return implementation.GetAccountMovementService(GetAccountMovementPort())
}

type AccountMovementRepository struct {
}

func (amr AccountMovementRepository) Deposit(accountId string) (string, error) {
	return "movementID-Generated", nil
}
