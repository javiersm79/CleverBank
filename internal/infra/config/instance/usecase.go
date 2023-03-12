package instance

import (
	"cleverbank/internal/core/implementation"
	"cleverbank/internal/core/usecase"
)

func GetAccountMovementUseCase() usecase.AccountMovementUseCase {
	return implementation.GetAccountMovementService(GetAccountMovementPort())
}

func GetAccountInfoUseCase() usecase.AccountMovementUseCase {
	return implementation.GetAccountInfoService(GetAccountInfoPort())
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
type AccountMovementRepository struct {
}

func (amr AccountMovementRepository) Deposit(accountId string) (string, error) {
	return "movementID-Generated", nil
}
