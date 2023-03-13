package controller

import (
	"cleverbank/internal/core/usecase"
	"sync"
)

var (
	instanceAccountMovementController *AccountMovementController
	onceAccountMovementController     sync.Once

	instanceAccountInfoController *AccountInfoController
	onceAccountInfoController     sync.Once

	instanceAccountCreationController *AccountCreationController
	onceAccountCreationController     sync.Once
)

func GetAccountMovementController(u usecase.AccountMovementUseCase) *AccountMovementController {
	onceAccountMovementController.Do(func() {
		instanceAccountMovementController = &AccountMovementController{u}
	})
	return instanceAccountMovementController
}

func GetAccountInfoController(u usecase.AccountInfoUseCase) *AccountInfoController {
	onceAccountInfoController.Do(func() {
		instanceAccountInfoController = &AccountInfoController{u}
	})
	return instanceAccountInfoController
}

func GetAccountCreationController(u usecase.AccountCreationUseCase) *AccountCreationController {
	onceAccountCreationController.Do(func() {
		instanceAccountCreationController = &AccountCreationController{u}
	})
	return instanceAccountCreationController
}
