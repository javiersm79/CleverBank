package instance

import (
	"cleverbank/internal/core/implementation"
	"cleverbank/internal/core/usecase"
)

func GetAccountMovementUseCase() usecase.AccountMovementUseCase {
	return implementation.GetAccountMovementService(GetAccountMovementPort())
}

func GetAccountInfoUseCase() usecase.AccountInfoUseCase {
	return implementation.GetAccountInfoService(GetAccountInfoPort())
}

func GetAccountCreationUseCase() usecase.AccountCreationUseCase {
	return implementation.GetAccountCreationService(GetAccountCreationPort())
}
