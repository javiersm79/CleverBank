package implementation

import (
	"cleverbank/internal/core/port"
	"sync"
)

var (
	instanceAccountMovementService *AccountMovementService
	onceAccountMovementService     sync.Once

	instanceAccountInfoService *AccountInfoService
	onceAccountInfoService     sync.Once

	instanceAccountCreationService *AccountCreationService
	onceAccountCreationService     sync.Once

	instanceAuthLoginService *AuthLoginService
	onceAuthLoginService     sync.Once
)

func GetAccountMovementService(am port.AccountMovementPort) *AccountMovementService {
	onceAccountMovementService.Do(func() {
		instanceAccountMovementService = &AccountMovementService{am}
	})
	return instanceAccountMovementService
}

func GetAccountInfoService(ai port.AccountInfoPort) *AccountInfoService {
	onceAccountInfoService.Do(func() {
		instanceAccountInfoService = &AccountInfoService{ai}
	})
	return instanceAccountInfoService
}

func GetAccountCreationService(ai port.AccountCreationPort) *AccountCreationService {
	onceAccountCreationService.Do(func() {
		instanceAccountCreationService = &AccountCreationService{ai}
	})
	return instanceAccountCreationService
}

func GetAuthLoginService(ai port.AuthLoginPort) *AuthLoginService {
	onceAuthLoginService.Do(func() {
		instanceAuthLoginService = &AuthLoginService{ai}
	})
	return instanceAuthLoginService
}
