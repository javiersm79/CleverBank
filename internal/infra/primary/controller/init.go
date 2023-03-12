package controller

import (
	"cleverbank/internal/core/usecase"
	"sync"
)

var (
	instanceAccountMovementController *AccountMovementController
	onceAccountMovementController     sync.Once
)

func GetAccountMovementController(u usecase.AccountMovementUseCase) *AccountMovementController {
	onceAccountMovementController.Do(func() {
		instanceAccountMovementController = &AccountMovementController{u}
	})
	return instanceAccountMovementController
}
