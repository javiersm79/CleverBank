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
