package implementation

import (
	"cleverbank/internal/core/port"
	"sync"
)

var (
	instanceAccountMovementService *AccountMovementService
	onceAccountMovementService     sync.Once
)

func GetAccountMovementService(am port.AccountMovementPort) *AccountMovementService {
	onceAccountMovementService.Do(func() {
		instanceAccountMovementService = &AccountMovementService{am}
	})
	return instanceAccountMovementService
}
